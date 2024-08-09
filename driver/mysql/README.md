# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/database/driver/mysql.svg)](https://pkg.go.dev/github.com/gopi-frame/database/driver/mysql)
[![Test driver mysql](https://github.com/gopi-frame/database/actions/workflows/mysql.yml/badge.svg)](https://github.com/gopi-frame/database/actions/workflows/mysql.yml)
[![codecov](https://codecov.io/gh/gopi-frame/database/graph/badge.svg?token=XRKDX3B3PN&flag=mysql)](https://codecov.io/gh/gopi-frame/database?flags[0]=mysql)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/database/driver/mysql)](https://goreportcard.com/report/github.com/gopi-frame/database/driver/mysql)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package mysql providers MySQL database driver for 
[`gopi-frame/database`](https://pkg.go.dev/gopi-frame/database) package.

## Installation
```shell
go get -u github.com/gopi-frame/database/driver/mysql
```

## Import
```go
import _ "github.com/gopi-frame/database/driver/mysql"
```

## Usage
```go
package main

import (
    "github.com/gopi-frame/database"

    _ "github.com/gopi-frame/database/driver/mysql"
)

func main() {
    db, err := database.Connect("mysql", map[string]any{
        "dsn": "user:password@tcp(127.0.0.1:3306)/database?parseTime=true",
    })
    if err!= nil {
        panic(err)
    }
}
```

## Options

This package uses package [mapstructure](https://pkg.go.dev/github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [mysql.Config](https://pkg.go.dev/gorm.io/driver/mysql#Config).

### Example
```go
var options = map[string]any{
    "dsn": "user:password@tcp(127.0.0.1:3306)/database?parseTime=true",
    "SkipInitializeWithVersion": true,
    "DefaultStringSize":         255,
}
```

