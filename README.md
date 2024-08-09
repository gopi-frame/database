# Database
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/database.svg)](https://pkg.go.dev/github.com/gopi-frame/database)
[![Go](https://github.com/gopi-frame/database/actions/workflows/go.yml/badge.svg)](https://github.com/gopi-frame/database/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gopi-frame/database/graph/badge.svg?token=N2LZNDNDCT&flag=database)](https://codecov.io/gh/gopi-frame/database?flags[0]=database)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/database)](https://goreportcard.com/report/github.com/gopi-frame/database)
[![Mit License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package database provides a database abstraction client.

This package is based on [gorm](https://github.com/go-gorm/gorm)

## Installation
```shell
go get -u github.com/gopi-frame/database
```

## Import
```go
import "github.com/gopi-frame/database"
```

## Usage

```go
package main

import (
    "github.com/gopi-frame/database"
    
    _ "github.com/gopi-frame/database/sqlite"
    // _ "github.com/gopi-frame/database/mysql"
    // _ "github.com/gopi-frame/database/postgres"
    // _ "github.com/gopi-frame/database/sqlserver"
)

func main() {
    db, err := database.Connect("sqlite", map[string]any{
        "dsn": "file:test.db",
    })
    if err!= nil {
        panic(err)
    }
}
```

## Drivers

- [sqlite](./driver/sqlite/README.md)
- [mysql](./driver/mysql/README.md)
- [postgres](./driver/postgres/README.md)
- [sqlserver](./driver/sqlserver/README.md)

## How to create a custom driver

To create a custom driver, just implement
the [database.Driver](https://pkg.go.dev/github.com/gopi-frame/contract/database#Driver) interface
and register it using [database.Register](https://pkg.go.dev/github.com/gopi-frame/database#Register).

### Example

```go
package main

import (
    "github.com/gopi-frame/database"
    "gorm.io/gorm"
)

var driverName = "custom"

func init() {
    database.Register(driverName, &CustomDriver{})
}

type CustomDriver struct{}

func (d *CustomDriver) Open(options map[string]any) (gorm.Dialector, error) {
    var d gorm.Dialector
    // implement your driver here
    return d, nil
}
```