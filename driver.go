package database

import (
	"fmt"

	"github.com/gopi-frame/collection/kv"
	"github.com/gopi-frame/database/driver"
	"github.com/gopi-frame/exception"
	"gorm.io/gorm"
)

var drivers = kv.NewMap[string, driver.Driver]()

func Register(driverName string, driver driver.Driver) {
	drivers.Lock()
	defer drivers.Unlock()
	if _, dup := drivers.Get(driverName); dup {
		panic(exception.NewArgumentException("driverName", driverName, fmt.Sprintf("duplicate driver \"%s\"", driverName)))
	}
	drivers.Set(driverName, driver)
}

func Open(driverName string, options map[string]any) (gorm.Dialector, error) {
	drivers.RLock()
	driver, ok := drivers.Get(driverName)
	drivers.RUnlock()
	if !ok {
		return nil, exception.NewArgumentException("driverName", driverName, fmt.Sprintf("Unknown driver \"%s\"", driverName))
	}
	return driver.Open(options)
}

func Connect(driverName string, options map[string]any, gormOpts ...gorm.Option) (*gorm.DB, error) {
	dialector, err := Open(driverName, options)
	if err != nil {
		return nil, err
	}
	return gorm.Open(dialector, gormOpts...)
}
