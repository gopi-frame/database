package sqlite

import (
	"bytes"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"net/url"
	"strings"
)

type Connector struct {
	DSN         string
	Database    string
	Params      map[string]string
	GormOptions *gorm.Config
	Replicas    []map[string]any
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
			database.ExpandStringWithEnvHookFunc(),
			database.ExpandSliceWithEnvHookFunc(),
			database.ExpandStringKeyMapWithEnvHookFunc(),
			mapstructure.StringToBasicTypeHookFunc(),
			database.NamingStrategyParseHookFunc(),
		),
	})
	if err != nil {
		panic(err)
	}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return connector, nil
}

func (c *Connector) buildDSN() string {
	var buf bytes.Buffer
	if len(c.Database) > 0 {
		buf.WriteString(c.Database)
	} else {
		buf.WriteString(":memory:")
	}
	if len(c.Params) > 0 {
		var params = make(url.Values)
		buf.WriteByte('?')
		for k, v := range c.Params {
			params.Set(k, v)
		}
		buf.WriteString(params.Encode())
	}
	return buf.String()
}

func (c *Connector) GetDSN() string {
	if len(c.DSN) > 0 {
		return c.DSN
	}
	return c.buildDSN()
}

func (c *Connector) Open() gorm.Dialector {
	return sqlite.Open(c.GetDSN())
}

func (c *Connector) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(c.Open(), c.GormOptions)
	if err != nil {
		return nil, err
	}
	var replicas = make([]gorm.Dialector, 0, len(c.Replicas))
	for _, replica := range c.Replicas {
		driver := replica["driver"].(string)
		replica, err := database.Open(driver, replica)
		if err != nil {
			return nil, err
		}
		replicas = append(replicas, replica)
	}
	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{c.Open()},
		Replicas: replicas,
	})); err != nil {
		return nil, err
	}
	return db, nil
}
