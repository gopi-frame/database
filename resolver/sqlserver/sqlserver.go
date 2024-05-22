package sqlserver

import (
	"fmt"
	"net/url"

	"github.com/gopi-frame/database"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Resolver resolver
type Resolver struct{}

// Resolve resolve
func (r *Resolver) Resolve(cfg database.ConnectionConfig) gorm.Dialector {
	dsnURL := url.URL{}
	dsnURL.Scheme = "sqlserver"
	if cfg.Password != "" {
		dsnURL.User = url.UserPassword(cfg.Username, cfg.Password)
	} else {
		dsnURL.User = url.User(cfg.Username)
	}
	if cfg.Port != 0 {
		dsnURL.Host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	} else {
		dsnURL.Host = cfg.Host
	}
	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}
	cfg.Params["database"] = cfg.Database
	query := make(url.Values)
	for key, value := range cfg.Params {
		query.Set(key, value)
	}
	dsnURL.RawQuery = query.Encode()
	dsn := dsnURL.String()
	return sqlserver.Open(dsn)
}
