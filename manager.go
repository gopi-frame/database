package database

import (
	"sync"

	"github.com/gopi-frame/collection/kv"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	once sync.Once
	*gorm.DB

	defaultConnection string
	connections       *kv.Map[string, *gorm.DB]
	lazyConnections   *kv.Map[string, func() (*gorm.DB, error)]
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		connections:     kv.NewMap[string, *gorm.DB](),
		lazyConnections: kv.NewMap[string, func() (*gorm.DB, error)](),
	}
}

func (m *DatabaseManager) SetDefaultConnection(connection string) {
	m.defaultConnection = connection
	m.once.Do(func() {
		m.DB = m.GetConnection(connection)
	})
}

func (m *DatabaseManager) AddConnection(name string, db *gorm.DB) {
	m.connections.Lock()
	defer m.connections.Unlock()
	m.connections.Set(name, db)
}

func (m *DatabaseManager) AddDeferConnection(name string, config map[string]any) {
	m.lazyConnections.Lock()
	defer m.lazyConnections.Unlock()
	m.lazyConnections.Set(name, func() (*gorm.DB, error) {
		driver := config["driver"].(string)
		return Connect(driver, config)
	})
}

func (m *DatabaseManager) HasConnection(name string) bool {
	m.connections.RLock()
	if m.connections.ContainsKey(name) {
		m.connections.RUnlock()
		return true
	}
	m.connections.RUnlock()
	m.lazyConnections.RLock()
	if m.lazyConnections.ContainsKey(name) {
		m.lazyConnections.RUnlock()
		return true
	}
	m.lazyConnections.RUnlock()
	return false
}

func (m *DatabaseManager) GetConnection(name string) *gorm.DB {
	m.connections.RLock()
	if conn, ok := m.connections.Get(name); ok {
		m.connections.RUnlock()
		return conn
	}
	m.connections.RUnlock()
	m.lazyConnections.RLock()
	if lazyConn, ok := m.lazyConnections.Get(name); ok {
		m.lazyConnections.RUnlock()
		m.connections.Lock()
		defer m.connections.Unlock()
		conn, err := lazyConn()
		if err != nil {
			panic(err)
		}
		m.connections.Set(name, conn)
		return conn
	}
	panic(NewNotConfiguredException(name))
}
