# Sqlite
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

This package uses package [mapstructure](github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [sqlite.Config](https://pkg.go.dev/gorm.io/driver/sqlite#Config).

### Example

```go
var options = map[string]any{
	"dsn": "file:test.db",
}
```