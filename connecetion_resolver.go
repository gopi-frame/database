package database

import "gorm.io/gorm"

// ConnectionResolver connection resolver interface
type ConnectionResolver interface {
	Resolve(any) gorm.Dialector
}
