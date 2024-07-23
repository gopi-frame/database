// Package mysql provides a mysql database driver.
package mysql

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/mysql"
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
// For more information on the options, see [mysql.Config](https://pkg.go.dev/gorm.io/driver/mysql#Config).
func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	var config mysql.Config
	err := mapstructure.WeakDecode(options, &config)
	if err != nil {
		return nil, err
	}
	return mysql.New(config), nil
}
