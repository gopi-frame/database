// Package database is a package for managing database drivers.
package database

import (
	"sort"

	"github.com/gopi-frame/collection/kv"
	"github.com/gopi-frame/contract/database"
	"github.com/gopi-frame/exception"
	"gorm.io/gorm"
)

var drivers = kv.NewMap[string, database.Driver]()

// Register registers a new database driver.
// If a driver with the same name already exists, it panics.
func Register(driverName string, driver database.Driver) {
	drivers.Lock()
	defer drivers.Unlock()
	if driver == nil {
		panic(exception.NewEmptyArgumentException("driver"))
	}
	if _, dup := drivers.Get(driverName); dup {
		panic(NewDuplicateDriverException(driverName))
	}
	drivers.Set(driverName, driver)
}

// Open opens a new database connector using the given driver name and options.
// If the driver with the given name doesn't exist, it panics.
func Open(driverName string, options map[string]any) (gorm.Dialector, error) {
	drivers.RLock()
	driver, ok := drivers.Get(driverName)
	drivers.RUnlock()
	if !ok {
		return nil, NewUnregisteredDriverException(driverName)
	}
	return driver.Open(options)
}

// Connect connects to a database using the given driver name and options.
func Connect(driverName string, options map[string]any) (*gorm.DB, error) {
	drivers.RLock()
	defer drivers.RUnlock()
	driver, ok := drivers.Get(driverName)
	if !ok {
		return nil, NewUnregisteredDriverException(driverName)
	}
	return driver.Connect(options)
}

// Drivers returns a list of registered database drivers.
func Drivers() []string {
	drivers.RLock()
	defer drivers.RUnlock()
	list := drivers.Keys()
	sort.Strings(list)
	return list
}
