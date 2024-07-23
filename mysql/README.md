# MySQL
Package mysql providers MySQL database driver.

## Installation
```shell
go get -u github.com/gopi-frame/database/mysql
```

## Import
```go
import _ "github.com/gopi-frame/database/mysql"
```

## Usage
```go
package main

import (
    "github.com/gopi-frame/database"

    _ "github.com/gopi-frame/database/mysql"
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

This package uses package [mapstructure](github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [mysql.Config](https://pkg.go.dev/gorm.io/driver/mysql#Config).

### Example
```go
var options = map[string]any{
    "dsn": "user:password@tcp(127.0.0.1:3306)/database?parseTime=true",
    "SkipInitializeWithVersion": true,
    "DefaultStringSize":         255,
}
```

