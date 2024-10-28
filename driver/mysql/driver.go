// Package mysql provides a mysql database driver.
package mysql

import (
	"github.com/gopi-frame/database"
	"gorm.io/gorm"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/database/mysql.driverName=custom`
var driverName = "mysql"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

// Driver is a mysql database driver.
type Driver struct{}

// Open opens a mysql database connector.
func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	connector, err := NewConnector(options)
	if err != nil {
		return nil, err
	}
	return connector.Open(), nil
}

func (Driver) Connect(options map[string]any) (*gorm.DB, error) {
	connector, err := NewConnector(options)
	if err != nil {
		return nil, err
	}
	return connector.Connect()
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (gorm.Dialector, error) {
	return new(Driver).Open(options)
}

func Connect(options map[string]any) (*gorm.DB, error) {
	return new(Driver).Connect(options)
}
