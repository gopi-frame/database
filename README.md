# Database
Package database is a package for managing database drivers.

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
    db, err := database.Open("sqlite", map[string]any{
        "dsn": "file:test.db",
    })
    if err!= nil {
        panic(err)
    }
}
```

## Drivers

- [sqlite](sqlite/README.md)
- [mysql](mysql/README.md)
- [postgres](postgres/README.md)
- [sqlserver](sqlserver/README.md)

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