package mysql

import (
	"crypto/tls"
	"fmt"
	mysqllib "github.com/go-sql-driver/mysql"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strings"
	"time"
)

type Connector struct {
	DriverName                    string
	ServerVersion                 string
	DSN                           string
	Username                      string
	Password                      string
	Net                           string
	Host                          string
	Port                          int
	Addr                          string
	Database                      string
	Params                        map[string]string
	Collation                     string
	Loc                           *time.Location
	MaxAllowedPacket              int
	ServerPubKey                  string
	TLSConfig                     string
	TLS                           *tls.Config
	Timeout                       time.Duration
	ReadTimeout                   time.Duration
	WriteTimeout                  time.Duration
	AllowAllFiles                 bool
	AllowCleartextPassword        bool
	AllowFallbackToPlaintext      bool
	AllowNativePasswords          bool
	AllowOldPasswords             bool
	CheckConnLiveness             bool
	ClientConnLiveness            bool
	ClientFoundRows               bool
	ColumnsWithAlias              bool
	InterpolateParams             bool
	MultiStatements               bool
	ParseTime                     bool
	RejectReadOnly                bool
	SkipInitializeWithVersion     bool
	DefaultStringSize             uint
	DefaultDatetimePrecision      *int
	DisableWithReturning          bool
	DisableDatetimePrecision      bool
	DontSupportRenameIndex        bool
	DontSupportRenameColumn       bool
	DontSupportForShareClause     bool
	DontSupportNullAsDefaultValue bool
	DontSupportRenameColumnUnique bool
	DontSupportDropConstraint     bool
	GormOptions                   *gorm.Config
	Replicas                      []map[string]any
}

func NewConnector(config map[string]any) (*Connector, error) {
	var connector = &Connector{
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
	}
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
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToBasicTypeHookFunc(),
			database.StringToLocationHookFunc(),
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

func (c *Connector) Open() gorm.Dialector {
	var cfg = mysql.Config{
		DriverName:                    c.DriverName,
		DSN:                           c.DSN,
		SkipInitializeWithVersion:     c.SkipInitializeWithVersion,
		DefaultStringSize:             c.DefaultStringSize,
		DefaultDatetimePrecision:      c.DefaultDatetimePrecision,
		DisableWithReturning:          c.DisableWithReturning,
		DisableDatetimePrecision:      c.DisableDatetimePrecision,
		DontSupportRenameIndex:        c.DontSupportRenameIndex,
		DontSupportRenameColumn:       c.DontSupportRenameColumn,
		DontSupportForShareClause:     c.DontSupportForShareClause,
		DontSupportNullAsDefaultValue: c.DontSupportNullAsDefaultValue,
		DontSupportRenameColumnUnique: c.DontSupportRenameColumnUnique,
		DontSupportDropConstraint:     c.DontSupportDropConstraint,
		DSNConfig: &mysqllib.Config{
			User:   c.Username,
			Passwd: c.Password,
			Net:    c.Net,
			Addr: func() string {
				if c.Addr != "" {
					return c.Addr
				}
				if c.Host == "" {
					c.Host = "localhost"
				}
				if c.Port <= 0 {
					c.Port = 3306
				}
				c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
				return c.Addr
			}(),
			DBName:                   c.Database,
			Params:                   c.Params,
			Collation:                c.Collation,
			Loc:                      c.Loc,
			MaxAllowedPacket:         c.MaxAllowedPacket,
			ServerPubKey:             c.ServerPubKey,
			TLSConfig:                c.TLSConfig,
			TLS:                      c.TLS,
			Timeout:                  c.Timeout,
			ReadTimeout:              c.ReadTimeout,
			WriteTimeout:             c.WriteTimeout,
			AllowAllFiles:            c.AllowAllFiles,
			AllowCleartextPasswords:  c.AllowCleartextPassword,
			AllowFallbackToPlaintext: c.AllowFallbackToPlaintext,
			AllowNativePasswords:     c.AllowNativePasswords,
			AllowOldPasswords:        c.AllowOldPasswords,
			CheckConnLiveness:        c.CheckConnLiveness,
			ClientFoundRows:          c.ClientFoundRows,
			ColumnsWithAlias:         c.ColumnsWithAlias,
			InterpolateParams:        c.InterpolateParams,
			MultiStatements:          c.MultiStatements,
			ParseTime:                c.ParseTime,
			RejectReadOnly:           c.RejectReadOnly,
		},
	}
	return mysql.New(cfg)
}

func (c *Connector) Connect() (*gorm.DB, error) {
	source := c.Open()
	var db *gorm.DB
	var err error
	if c.GormOptions == nil {
		db, err = gorm.Open(source)
	} else {
		db, err = gorm.Open(source, c.GormOptions)
	}
	if err != nil {
		return nil, err
	}
	if len(c.Replicas) == 0 {
		return db, nil
	}
	replicas := make([]gorm.Dialector, 0, len(c.Replicas))
	for _, replica := range c.Replicas {
		driverName := replica["driver"].(string)
		replica, err := database.Open(driverName, replica)
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
