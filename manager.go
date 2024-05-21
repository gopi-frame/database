package database

import (
	"github.com/gopi-frame/support/maps"
	"gorm.io/gorm"
)

// NewManager new manager
func NewManager(defaultConnection string, factory *ConnectionFactory) *Manager {
	manager := new(Manager)
	manager.defaultConnection = defaultConnection
	manager.connections = maps.NewMap[string, *Connection]()
	manager.factory = factory
	return manager
}

// Manager manager
type Manager struct {
	defaultConnection string
	*Connection       // default connection
	connections       *maps.Map[string, *Connection]
	factory           *ConnectionFactory
}

// SetDefaultConnection set default connection
func (m *Manager) SetDefaultConnection(name string) {
	m.defaultConnection = name
}

// DB get connection instance
func (m *Manager) DB(name string) *Connection {
	connection, ok := m.connections.Get(name)
	if ok {
		connection.lazyInit()
		return connection
	}
	panic(NewConnectionNotFoundException(name))
}

// AddConnection add a new connection
func (m *Manager) AddConnection(name string, db *gorm.DB) {
	m.connections.Set(name, NewConnection(db))
}

// AddLazyConnection add a new lazy connection
func (m *Manager) AddLazyConnection(name string, connector func() (*gorm.DB, error)) {
	m.connections.Set(name, NewLazyConnection(connector))
}

// RegisterResolver register resolver
func (m *Manager) RegisterResolver(name string, resolver ConnectionResolver) {
	m.factory.RegisterResolver(name, resolver)
}

// Resolve resolve connection from config
func (m *Manager) Resolve(config ConnectionConfig) {
	m.connections.Set(config.Name, m.factory.Resolve(config))
}
