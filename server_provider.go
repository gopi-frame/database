package database

import (
	"reflect"

	"github.com/gopi-frame/contract/config"
	"github.com/gopi-frame/contract/container"
	"github.com/gopi-frame/contract/support"
)

// ServerProvider database server provider
type ServerProvider struct {
	support.ServerProvider `json:"-" toml:"-" yaml:"-" mapstructure:"-"`
	Default                string             `json:"default" toml:"default" yaml:"default" mapstructure:"default"`
	Connections            []ConnectionConfig `json:"connections" toml:"connections" yaml:"connections" mapstructure:"connections"`
}

// Register register
func (d *ServerProvider) Register(c container.Container) {
	c.Bind("database", func(c container.Container) any {
		c.Get("config").(config.Repository).Unmarshal("database", d)
		return NewManager(d.Default, NewConnectionFactory())
	})
	c.Alias("database", "db")
	c.Alias("database", reflect.TypeFor[Manager]().String())
	c.Alias("database", reflect.TypeFor[Connection]().String())
}

// Boot boot
func (d *ServerProvider) Boot(c container.Container) {
	manager := c.Get("db").(*Manager)
	for _, config := range d.Connections {
		manager.Resolve(config)
	}
}
