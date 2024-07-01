package sqlserver

import (
	"strings"

	"github.com/gopi-frame/database"
	"github.com/gopi-frame/exception"
	"github.com/gopi-frame/utils/kv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var driverName = "sqlserver"

func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	config := sqlserver.Config{}
	dsn, err := kv.GetE[string](options, OptKeyDSN)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(dsn) == "" {
		return nil, exception.New("dsn can't be empty")
	}
	config.DSN = dsn
	defaultStringSize := kv.Get[int](options, OptKeyDefaultStringSize)
	config.DefaultStringSize = defaultStringSize
	return sqlserver.New(config), nil
}
