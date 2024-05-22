package mysql

import (
	"strconv"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/gopi-frame/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Resolver mysql connection resolver
type Resolver struct{}

// Resolve resolve
func (r *Resolver) Resolve(cfg database.ConnectionConfig) gorm.Dialector {
	address := cfg.Host
	if cfg.Port > 0 {
		address += ":" + strconv.Itoa(cfg.Port)
	}
	mysqlConfig := mysql.Config{
		DSNConfig: &mysqldriver.Config{
			User:      cfg.Username,
			Passwd:    cfg.Password,
			Net:       cfg.Protocol,
			Addr:      address,
			DBName:    cfg.Database,
			Params:    cfg.Params,
			Collation: cfg.Collation,
			Loc: func() *time.Location {
				location, err := time.LoadLocation(cfg.Location)
				if err != nil {
					panic(err)
				}
				return location
			}(),
			ParseTime: cfg.ParseTime,
		},
	}
	return mysql.New(mysqlConfig)
}
