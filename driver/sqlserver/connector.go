package sqlserver

import (
	"bytes"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"github.com/gopi-frame/env"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Connector struct {
	DriverName        string
	DSN               string
	Host              string
	Port              int
	Username          string
	Password          string
	Database          string
	Charset           string
	Params            map[string]string
	GormOptions       *gorm.Config
	Replicas          []map[string]any
	DefaultStringSize int
}

func NewConnector(config map[string]any) (*Connector, error) {
	var connector = new(Connector)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           connector,
		WeaklyTypedInput: true,
		MatchName: func(mapKey, fieldName string) bool {
			return strings.EqualFold(mapKey, fieldName) ||
				strings.EqualFold(fieldName, strings.ReplaceAll(mapKey, "_", ""))
		},
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			env.ExpandStringWithEnvHookFunc(),
			env.ExpandSliceWithEnvHookFunc(),
			env.ExpandStringKeyMapWithEnvHookFunc(),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToBasicTypeHookFunc(),
			database.NamingStrategyParseHookFunc(),
		),
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return connector, nil
}

func (c *Connector) buildDSN() string {
	var buf bytes.Buffer
	buf.WriteString("sqlserver://")
	if len(c.Username) > 0 {
		buf.WriteString(c.Username)
		if len(c.Password) > 0 {
			buf.WriteByte(':')
			buf.WriteString(c.Password)
		}
		buf.WriteByte('@')
	}
	if len(c.Host) > 0 {
		buf.WriteString(c.Host)
		if c.Port > 0 {
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(c.Port))
		}
	}
	buf.WriteByte('/')
	buf.WriteString(c.Database)
	params := make(url.Values)
	if len(c.Charset) > 0 {
		params.Set("charset", c.Charset)
	}
	if len(c.Params) > 0 {
		for k, v := range c.Params {
			params.Set(k, v)
		}
	}
	if len(params) > 0 {
		buf.WriteByte('?')
		buf.WriteString(params.Encode())
	}
	return buf.String()
}

func (c *Connector) GetDSN() string {
	if c.DSN != "" {
		return c.DSN
	}
	return c.buildDSN()
}

func (c *Connector) Open() gorm.Dialector {
	return sqlserver.Open(c.GetDSN())
}

func (c *Connector) Connect() (*gorm.DB, error) {
	source := c.Open()
	db, err := gorm.Open(source, c.GormOptions)
	if err != nil {
		return nil, err
	}
	if len(c.Replicas) == 0 {
		return db, nil
	}
	var replicas []gorm.Dialector
	for _, replica := range c.Replicas {
		driver := replica["driver"].(string)
		replica, err := database.Open(driver, replica)
		if err != nil {
			return nil, err
		}
		replicas = append(replicas, replica)
	}
	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{source},
		Replicas: replicas,
	})); err != nil {
		return nil, err
	}
	return db, nil
}
