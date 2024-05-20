package database

import (
	"reflect"

	"github.com/gopi-frame/contract/container"
	"github.com/gopi-frame/contract/support"
)

// ServerProvider database server provider
type ServerProvider struct {
	support.ServerProvider
}

// Register register
func (d *ServerProvider) Register(c container.Container) {
	c.Bind("database", func(c container.Container) any {
		return NewManager()
	})
	c.Alias("database", "db")
	c.Alias("database", reflect.TypeFor[Manager]().String())
	c.Alias("database", reflect.TypeFor[Connection]().String())
}
