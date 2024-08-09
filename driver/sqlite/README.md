# Sqlite
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/database/driver/sqlite.svg)](https://pkg.go.dev/github.com/gopi-frame/database/driver/sqlite)
[![Test driver sqlite](https://github.com/gopi-frame/database/actions/workflows/sqlite.yml/badge.svg)](https://github.com/gopi-frame/database/actions/workflows/sqlite.yml)
[![codecov](https://codecov.io/gh/gopi-frame/database/graph/badge.svg?token=XRKDX3B3PN&flag=sqlite)](https://codecov.io/gh/gopi-frame/database?flags[0]=sqlite)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/database/driver/sqlite)](https://goreportcard.com/report/github.com/gopi-frame/database/driver/sqlite)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package sqlite provides sqlite database driver.

## Installation
```shell
go get -u github.com/gopi-frame/database/sqlite
```

## Import
```go
import _ "github.com/gopi-frame/database/sqlite"
```

## Usage
```go
package main

import (
	"github.com/gopi-frame/database"
	
	_ "github.com/gopi-frame/database/sqlite"
)

func main() {
	db, err := database.Connect("sqlite", map[string]any{
		"dsn": "file:test.db",
    })
	if err != nil {
		panic(err)
	}
}
```

## Options

This package uses package [mapstructure](https://github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [sqlite.Config](https://pkg.go.dev/gorm.io/driver/sqlite#Config).

### Example

```go
var options = map[string]any{
	"dsn": "file:test.db",
}
```