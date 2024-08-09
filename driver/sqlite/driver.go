// Package sqlite provides sqlite database driver.
package sqlite

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/sqlite"
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
// For more information on the options, see [sqlite.Config](https://pkg.go.dev/gorm.io/driver/sqlite#Config).
func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	var config sqlite.Config
	err := mapstructure.WeakDecode(options, &config)
	if err != nil {
		return nil, err
	}
	return sqlite.New(config), nil
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (gorm.Dialector, error) {
	return new(Driver).Open(options)
}
