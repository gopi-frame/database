package sqlite

import (
	"github.com/gopi-frame/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Resolver sqlite connection resolver
type Resolver struct{}

// Resolve resolve
func (r *Resolver) Resolve(cfg database.ConnectionConfig) gorm.Dialector {
	return sqlite.Open(cfg.Database)
}
