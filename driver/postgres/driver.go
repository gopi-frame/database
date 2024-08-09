// Package postgres provides postgres database driver.
package postgres

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/database/postgres.driverName=custom`
var driverName = "postgres"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

// Driver is a postgres database driver.
type Driver struct{}

// Open opens a postgres database connector.
// For more information on the options, see [postgres.Config](https://pkg.go.dev/gorm.io/driver/postgres#Config).
func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	var config postgres.Config
	err := mapstructure.WeakDecode(options, &config)
	if err != nil {
		return nil, err
	}
	return postgres.New(config), nil
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (gorm.Dialector, error) {
	return new(Driver).Open(options)
}
