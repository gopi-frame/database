# Sqlserver
Package sqlserver providers sqlserver database driver.

## Installation

```shell
go get -u github.com/gopi-frame/database/sqlserver
```

## Import

```go
import _ "github.com/gopi-frame/database/sqlserver"
```

## Usage

```go
package main

import (
	"github.com/gopi-frame/database"
	
	_ "github.com/gopi-frame/database/sqlserver"
)

func main() {
	db, err := database.Connect("sqlserver", map[string]any{
		"dsn": "sqlserver://user:password@localhost:1433?database=test",
    })
	if err!= nil {
		panic(err)
	}
}
```

## Options

This package uses package [mapstructure](github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [sqlserver.Config](https://pkg.go.dev/gorm.io/driver/sqlserver#Config).

### Example

```go
var options = map[string]any{
	"dsn": "sqlserver://user:password@localhost:1433?database=test",
	"DefaultStringSize": 255,
}
```