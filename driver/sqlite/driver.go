// Package sqlite provides sqlite database driver.
package sqlite

import (
	"github.com/gopi-frame/database"
	"gorm.io/gorm"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/database/sqlite.driverName=custom`
var driverName = "sqlite"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

// Driver is a sqlite database driver.
type Driver struct{}

// Open opens a sqlite database connector.
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
