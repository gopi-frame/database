package postgres

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/gopi-frame/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Resolver postgres connection resolver
type Resolver struct{}

// Resolve resolve
func (r *Resolver) Resolve(cfg database.ConnectionConfig) gorm.Dialector {
	builder := new(strings.Builder)
	builder.WriteString("postgres://")
	builder.WriteString(cfg.Username)
	if cfg.Password != "" {
		builder.WriteString(":")
		builder.WriteString(cfg.Password)
	}
	builder.WriteString(cfg.Host)
	if cfg.Port > 0 {
		builder.WriteString(strconv.Itoa(cfg.Port))
	}
	builder.WriteByte('/')
	builder.WriteString(cfg.Database)
	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}
	if len(cfg.Params) > 0 {
		params := make(url.Values)
		for key, value := range cfg.Params {
			params.Set(key, value)
		}
		builder.WriteByte('?')
		builder.WriteString(params.Encode())
	}
	dsn := builder.String()
	return postgres.Open(dsn)
}
