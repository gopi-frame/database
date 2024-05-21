package database

// ConnectionConfig mysql connection config
type ConnectionConfig struct {
	Name     string             `json:"name" toml:"name" yaml:"name" mapstructure:"name"`
	Driver   string             `json:"driver" toml:"driver" yaml:"driver" mapstructure:"driver"`
	Host     string             `json:"host" toml:"host" yaml:"host" mapstructure:"host"`
	Port     int                `json:"port" toml:"port" yaml:"port" mapstructure:"port"`
	Database string             `json:"database" toml:"database" yaml:"database" mapstructure:"database"`
	Prefix   string             `json:"prefix" toml:"prefix" yaml:"prefix" mapstructure:"prefix"`
	Username string             `json:"username" toml:"username" yaml:"username" mapstructure:"username"`
	Password string             `json:"password" toml:"password" yaml:"password" mapstructure:"password"`
	Params   map[string]string  `json:"params" toml:"params" yaml:"params" mapstructure:"params"`
	Slaves   []ConnectionConfig `json:"slaves" toml:"slaves" yaml:"slaves" mapstructure:"slaves"`
	// below fields are for mysql only
	Protocol  string `json:"protocol" toml:"protocol" yaml:"protocol" mapstructure:"protocol"`
	Charset   string `json:"charset" toml:"charset" yaml:"charset" mapstructure:"charset"`
	Collation string `json:"collation" toml:"collation" yaml:"collation" mapstructure:"collation"`
	ParseTime bool   `json:"parseTime" toml:"parseTime" yaml:"parseTime" mapstructure:"parseTime"`
	Location  string `json:"location" toml:"location" yaml:"location" mapstructure:"location"`
}
