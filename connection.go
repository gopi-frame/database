package database

import (
	"sync"

	"gorm.io/gorm"
)

// NewConnection new connection
func NewConnection(db *gorm.DB) *Connection {
	return &Connection{
		DB: db,
	}
}

// NewLazyConnection new connection
func NewLazyConnection(connector func() (*gorm.DB, error)) *Connection {
	conn := new(Connection)
	conn.connector = sync.OnceValue[*gorm.DB](func() *gorm.DB {
		db, err := connector()
		if err != nil {
			panic(err)
		}
		return db
	})
	return conn
}

// Connection connection
type Connection struct {
	*gorm.DB
	connector func() *gorm.DB
}

func (c *Connection) lazyInit() {
	if c.connector != nil && c.DB == nil {
		c.DB = c.connector()
	}
}
