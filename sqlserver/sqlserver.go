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
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}
	cfg.Params["database"] = cfg.Database
	params := make(url.Values)
	for key, value := range cfg.Params {
		params.Set(key, value)
	}
	dsn = fmt.Sprintf("%s?%s", dsn, params.Encode())
	return sqlserver.Open(dsn)
}
