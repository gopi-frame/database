package postgres

import (
	"strings"

	"github.com/gopi-frame/database"
	"github.com/gopi-frame/exception"
	"github.com/gopi-frame/util/kv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var driverName = "postgres"

func init() {
	if driverName != "" {
		database.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (Driver) Open(options map[string]any) (gorm.Dialector, error) {
	config := postgres.Config{}
	dsn, err := kv.GetE[string](options, OptKeyDSN)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(dsn) == "" {
		return nil, exception.New("dsn can't be empty")
	}
	config.WithoutQuotingCheck = kv.Get[bool](options, OptKeyWithoutQuatingCheck)
	config.PreferSimpleProtocol = kv.Get[bool](options, OptKeyPreferSimpleProtocol)
	config.WithoutReturning = kv.Get[bool](options, OptKeyWithoutReturning)
	return postgres.New(config), nil
}
