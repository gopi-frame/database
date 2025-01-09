package postgres

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/database"
	"github.com/gopi-frame/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Connector struct {
	DriverName              string
	DSN                     string
	Host                    string
	Port                    int
	HostAddr                []string
	Username                string
	Password                string
	Passfile                string
	Database                string
	RequireAuth             string
	ChannelBinding          string
	ConnectTimeout          time.Duration
	SSLMode                 string
	Timezone                string
	ClientEncoding          string
	Options                 string
	ApplicationName         string
	FallbackApplicationName string
	KeepAlives              bool
	KeepAlivesIdle          time.Duration
	KeepAlivesInterval      time.Duration
	KeepAlivesCount         int
	TCPUserTimeout          time.Duration
	Replication             string
	GSSEncMode              string
	SSLNegotiation          string
	SSLCompression          bool
	SSLCert                 string
	SSLKey                  string
	SSLPassword             string
	SSLCertMode             string
	SSLRootCert             string
	SSLCrl                  string
	SSLCrlDir               string
	SSLSni                  bool
	RequirePeer             string
	SSLMinProtocolVersion   string
	SSLMaxProtocolVersion   string
	KrbSrvName              string
	GSSLib                  string
	GSSDelegation           bool
	Service                 string
	TargetSessionAttrs      string
	LoadBalanceHosts        string
	WithoutQuotingCheck     bool
	PreferSimpleProtocol    bool
	WithoutReturning        bool
	GormOptions             *gorm.Config
	Replicas                []map[string]any
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
			mapstructure.StringToSliceHookFunc(","),
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

func (c *Connector) buildDSN() string {
	var buf bytes.Buffer
	if c.Host != "" {
		buf.WriteString("host=" + c.Host)
		buf.WriteByte(' ')
	}
	if c.Port != 0 {
		buf.WriteString("port=" + strconv.Itoa(c.Port))
		buf.WriteByte(' ')
	}
	if len(c.HostAddr) > 0 {
		buf.WriteString("hosts=" + strings.Join(c.HostAddr, ","))
		buf.WriteByte(' ')
	}
	if c.Username != "" {
		buf.WriteString("user=" + c.Username)
		buf.WriteByte(' ')
	}
	if c.Password != "" {
		buf.WriteString("password=" + c.Password)
		buf.WriteByte(' ')
	}
	if c.Passfile != "" {
		buf.WriteString("passfile=" + c.Passfile)
		buf.WriteByte(' ')
	}
	if c.RequireAuth != "" {
		buf.WriteString("require_auth=" + c.RequireAuth)
		buf.WriteByte(' ')
	}
	if c.Database != "" {
		buf.WriteString("dbname=" + c.Database)
		buf.WriteByte(' ')
	}
	if c.ChannelBinding != "" {
		buf.WriteString("channel_binding=" + c.ChannelBinding)
		buf.WriteByte(' ')
	}
	if c.ConnectTimeout > 0 {
		buf.WriteString("connect_timeout=" + strconv.Itoa(int(c.ConnectTimeout.Seconds())))
		buf.WriteByte(' ')
	}
	if len(c.SSLMode) > 0 {
		buf.WriteString("sslmode=" + c.SSLMode)
		buf.WriteByte(' ')
	}
	if c.Timezone != "" {
		buf.WriteString("timezone=" + c.Timezone)
		buf.WriteByte(' ')
	}
	if c.ClientEncoding != "" {
		buf.WriteString("client_encoding=" + c.ClientEncoding)
		buf.WriteByte(' ')
	}
	if c.Options != "" {
		buf.WriteString("options=" + c.Options)
		buf.WriteByte(' ')
	}
	if c.ApplicationName != "" {
		buf.WriteString("application_name=" + c.ApplicationName)
		buf.WriteByte(' ')
	}
	if c.FallbackApplicationName != "" {
		buf.WriteString("fallback_application_name=" + c.FallbackApplicationName)
		buf.WriteByte(' ')
	}
	if !c.KeepAlives {
		buf.WriteString("keepalive=0")
		buf.WriteByte(' ')
	}
	if c.KeepAlivesIdle > 0 {
		buf.WriteString("keepalives_idle=" + strconv.Itoa(int(c.KeepAlivesIdle.Seconds())))
		buf.WriteByte(' ')
	}
	if c.KeepAlivesInterval > 0 {
		buf.WriteString("keepalives_interval=" + strconv.Itoa(int(c.KeepAlivesInterval.Seconds())))
		buf.WriteByte(' ')
	}
	if c.KeepAlivesCount > 0 {
		buf.WriteString("keepalives_count=" + strconv.Itoa(c.KeepAlivesCount))
		buf.WriteByte(' ')
	}
	if c.TCPUserTimeout > 0 {
		buf.WriteString("tcp_user_timeout=" + strconv.Itoa(int(c.TCPUserTimeout.Seconds())))
		buf.WriteByte(' ')
	}
	if c.Replication != "" {
		buf.WriteString("replication=" + c.Replication)
		buf.WriteByte(' ')
	}
	if c.GSSEncMode != "" {
		buf.WriteString("gssencmode=" + c.GSSEncMode)
		buf.WriteByte(' ')
	}
	if c.SSLNegotiation != "" {
		buf.WriteString("sslnegotiation=" + c.SSLNegotiation)
		buf.WriteByte(' ')
	}
	if c.SSLCompression {
		buf.WriteString("sslcompression=1")
		buf.WriteByte(' ')
	}
	if c.SSLCert != "" {
		buf.WriteString("sslcert=" + c.SSLCert)
		buf.WriteByte(' ')
	}
	if c.SSLKey != "" {
		buf.WriteString("sslkey=" + c.SSLKey)
		buf.WriteByte(' ')
	}
	if c.SSLPassword != "" {
		buf.WriteString("sslpassword=" + c.SSLPassword)
		buf.WriteByte(' ')
	}
	if c.SSLCertMode != "" {
		buf.WriteString("sslcertmode=" + c.SSLCertMode)
		buf.WriteByte(' ')
	}
	if c.SSLRootCert != "" {
		buf.WriteString("sslrootcert=" + c.SSLRootCert)
		buf.WriteByte(' ')
	}
	if c.SSLCrl != "" {
		buf.WriteString("sslcrl=" + c.SSLCrl)
		buf.WriteByte(' ')
	}
	if c.SSLCrlDir != "" {
		buf.WriteString("sslcrldir=" + c.SSLCrlDir)
		buf.WriteByte(' ')
	}
	if !c.SSLSni {
		buf.WriteString("sslsni=0")
		buf.WriteByte(' ')
	}
	if c.RequirePeer != "" {
		buf.WriteString("requirepeer=" + c.RequirePeer)
		buf.WriteByte(' ')
	}
	if c.SSLMinProtocolVersion != "" {
		buf.WriteString("ssl_min_protocol_version=" + c.SSLMinProtocolVersion)
		buf.WriteByte(' ')
	}
	if c.SSLMaxProtocolVersion != "" {
		buf.WriteString("ssl_max_protocol_version=" + c.SSLMaxProtocolVersion)
		buf.WriteByte(' ')
	}
	if c.KrbSrvName != "" {
		buf.WriteString("krbsrvname=" + c.KrbSrvName)
		buf.WriteByte(' ')
	}
	if c.GSSLib != "" {
		buf.WriteString("gsslib=" + c.GSSLib)
		buf.WriteByte(' ')
	}
	if c.GSSDelegation {
		buf.WriteString("gssdelegation=1")
		buf.WriteByte(' ')
	}
	if c.Service != "" {
		buf.WriteString("service=" + c.Service)
		buf.WriteByte(' ')
	}
	if c.TargetSessionAttrs != "" {
		buf.WriteString("target_session_attrs=" + c.TargetSessionAttrs)
		buf.WriteByte(' ')
	}
	if c.LoadBalanceHosts != "" {
		buf.WriteString("load_balance_hosts=" + c.LoadBalanceHosts)
		buf.WriteByte(' ')
	}
	c.DSN = strings.TrimSpace(buf.String())
	return c.DSN
}

// GetDSN returns the DSN.
func (c *Connector) GetDSN() string {
	if len(c.DSN) > 0 {
		return c.DSN
	}
	return c.buildDSN()
}

func (c *Connector) Open() gorm.Dialector {
	return postgres.Open(c.GetDSN())
}

func (c *Connector) Connect() (*gorm.DB, error) {
	source := c.Open()
	db, err := gorm.Open(source, c.GormOptions)
	if err != nil {
		return nil, err
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
