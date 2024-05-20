package database

import (
	"github.com/gopi-frame/support/maps"
	"gorm.io/gorm"
)

var defaultConnection = "default"

// NewManager new manager
func NewManager() *Manager {
	manager := new(Manager)
	manager.defaultConnection = defaultConnection
	manager.connections = maps.NewMap[string, *Connection]()
	return manager
}

// Manager manager
type Manager struct {
	defaultConnection string
	*Connection       // default connection
	connections       *maps.Map[string, *Connection]
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
