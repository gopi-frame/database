package database

import "gorm.io/gorm/logger"

// ConnectionConfig mysql connection config
type ConnectionConfig struct {
	Name     string             `json:"name" toml:"name" yaml:"name" mapstructure:"name"`
	Driver   string             `json:"driver" toml:"driver" yaml:"driver" mapstructure:"driver"`
	Host     string             `json:"host" toml:"host" yaml:"host" mapstructure:"host"`
	Port     int                `json:"port" toml:"port" yaml:"port" mapstructure:"port"`
	Database string             `json:"database" toml:"database" yaml:"database" mapstructure:"database"`
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
	// below fields are for gorm.Config
	Prefix                                   string          `json:"prefix" toml:"prefix" yaml:"prefix" mapstructure:"prefix"`
	SingularTable                            bool            `json:"singluarTable" toml:"singluarTable" yaml:"singluarTable" mapstructure:"singluarTable"`
	SkipDefaultTransaction                   bool            `json:"skipDefaultTransaction" toml:"skipDefaultTransaction" yaml:"skipDefaultTransaction" mapstructure:"skipDefaultTransaction"`
	FullSaveAssociations                     bool            `json:"fullSaveAssociations" toml:"fullSaveAssociations" yaml:"fullSaveAssociations" mapstructure:"fullSaveAssociations"`
	LoggerLevel                              logger.LogLevel `json:"loggerLevel" toml:"loggerLevel" yaml:"loggerLevel" mapstructure:"loggerLevel"`
	DryRun                                   bool            `json:"dryRun" toml:"dryRun" yaml:"dryRun" mapstructure:"dryRun"`
	PrepareStmt                              bool            `json:"prepareStmt" toml:"prepareStmt" yaml:"prepareStmt" mapstructure:"prepareStmt"`
	DisableAutomaticPing                     bool            `json:"disableAutomaticPing" toml:"disableAutomaticPing" yaml:"disableAutomaticPing" mapstructure:"disableAutomaticPing"`
	DisableForeignKeyConstraintWhenMigrating bool            `json:"disableForeignKeyConstraintWhenMigrating" toml:"disableForeignKeyConstraintWhenMigrating" yaml:"disableForeignKeyConstraintWhenMigrating" mapstructure:"disableForeignKeyConstraintWhenMigrating"`
	IgnoreRelationshipsWhenMigrating         bool            `json:"ignoreRelationshipsWhenMigrating" toml:"ignoreRelationshipsWhenMigrating" yaml:"ignoreRelationshipsWhenMigrating" mapstructure:"ignoreRelationshipsWhenMigrating"`
	DisableNestedTransaction                 bool            `json:"disableNestedTransaction" toml:"disableNestedTransaction" yaml:"disableNestedTransaction" mapstructure:"disableNestedTransaction"`
	AllowGlobalUpdate                        bool            `json:"allowGlobalUpdate" toml:"allowGlobalUpdate" yaml:"allowGlobalUpdate" mapstructure:"allowGlobalUpdate"`
	QueryFields                              bool            `json:"queryFields" toml:"queryFields" yaml:"queryFields" mapstructure:"queryFields"`
	CreateBatchSize                          int             `json:"createBatchSize" toml:"createBatchSize" yaml:"createBatchSize" mapstructure:"createBatchSize"`
	TranslateError                           bool            `json:"translateError" toml:"translateError" yaml:"translateError" mapstructure:"translateError"`
}
