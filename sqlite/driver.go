package sqlite

import (
	"strings"

	"github.com/gopi-frame/database"
	"github.com/gopi-frame/exception"
	"github.com/gopi-frame/util/kv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var driverName = "sqlite"

func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	dsn, err := kv.GetE[string](options, OptKeyDSN)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(dsn) == "" {
		return nil, exception.New("dsn can't be empty")
	}
	return sqlite.Open(dsn), nil
}
