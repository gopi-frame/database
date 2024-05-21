package postgres

import (
	"fmt"
	"net/url"

	"github.com/gopi-frame/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Resolver postgres connection resolver
type Resolver struct{}

// Resolve resolve
func (r *Resolver) Resolve(cfg database.ConnectionConfig) gorm.Dialector {
	dsnURL := url.URL{}
	dsnURL.Scheme = "postgres"
	if cfg.Password == "" {
		dsnURL.User = url.User(cfg.Username)
	} else {
		dsnURL.User = url.UserPassword(cfg.Username, cfg.Password)
	}
	if cfg.Port != 0 {
		dsnURL.Host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	} else {
		dsnURL.Host = cfg.Host
	}
	dsnURL.Path = cfg.Database
	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}
	if len(cfg.Params) > 0 {
		query := make(url.Values)
		for key, value := range cfg.Params {
			query.Set(key, value)
		}
		dsnURL.RawQuery = query.Encode()
	}
	dsn := dsnURL.String()
	return postgres.Open(dsn)
}
