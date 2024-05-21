package database

import (
	"fmt"

	"github.com/gopi-frame/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// NewConnectionFactory new connection factory
func NewConnectionFactory() *ConnectionFactory {
	return &ConnectionFactory{
		resolvers: map[string]ConnectionResolver{},
	}
}

// ConnectionFactory connection factory
type ConnectionFactory struct {
	resolvers map[string]ConnectionResolver
}

// Resolve resolve
func (c *ConnectionFactory) Resolve(config ConnectionConfig) *Connection {
	if len(config.Slaves) == 0 {
		return c.createSingleConnection(config)
	}
	return c.createReadAndWriteConnection(config)
}

// RegisterResolver register resolver
func (c *ConnectionFactory) RegisterResolver(name string, resolver ConnectionResolver) {
	c.resolvers[name] = resolver
}

func (c *ConnectionFactory) resolve(config ConnectionConfig) gorm.Dialector {
	var dialector gorm.Dialector
	resolver, ok := c.resolvers[config.Driver]
	if !ok {
		panic(exception.NewUnsupportedException(fmt.Sprintf("driver \"%s\" is not supported", config.Driver)))
	}
	dialector = resolver.Resolve(config)
	return dialector
}

func (c *ConnectionFactory) createSingleConnection(config ConnectionConfig) *Connection {
	return NewLazyConnection(func() (*gorm.DB, error) {
		return gorm.Open(c.resolve(config))
	})
}

func (c *ConnectionFactory) createReadAndWriteConnection(config ConnectionConfig) *Connection {
	master := c.resolve(config)
	slaves := []gorm.Dialector{}
	for _, readConfig := range config.Slaves {
		slaves = append(slaves, c.resolve(readConfig))
	}
	return NewLazyConnection(func() (*gorm.DB, error) {
		db, err := gorm.Open(master, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   config.Prefix,
				SingularTable: config.SingularTable,
			},
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			return nil, err
		}
		if err := db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: slaves,
		})); err != nil {
			return nil, err
		}
		return db, err
	})
}
