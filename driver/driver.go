package driver

import "gorm.io/gorm"

type Driver interface {
	Open(options map[string]any) (gorm.Dialector, error)
}
