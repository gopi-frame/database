# Sqlserver
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/database/driver/sqlserver.svg)](https://pkg.go.dev/github.com/gopi-frame/database/driver/sqlserver)
[![Test driver sqlserver](https://github.com/gopi-frame/database/actions/workflows/sqlserver.yml/badge.svg)](https://github.com/gopi-frame/database/actions/workflows/sqlserver.yml)
[![codecov](https://codecov.io/gh/gopi-frame/database/graph/badge.svg?token=XRKDX3B3PN&flag=sqlserver)](https://codecov.io/gh/gopi-frame/database?flag=sqlserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/database/driver/sqlserver)](https://goreportcard.com/report/github.com/gopi-frame/database/driver/sqlserver)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

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

This package uses package [mapstructure](https://github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [sqlserver.Config](https://pkg.go.dev/gorm.io/driver/sqlserver#Config).

### Example

```go
var options = map[string]any{
	"dsn": "sqlserver://user:password@localhost:1433?database=test",
	"DefaultStringSize": 255,
}
```