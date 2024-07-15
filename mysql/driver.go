package mysql

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var driverName = "mysql"

func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	config := new(mysql.Config)
	err := mapstructure.WeakDecode(options, config)
	if err != nil {
		return nil, err
	}
	return mysql.New(*config), nil
}
