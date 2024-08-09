# Postgres
Package postgres provides postgres database driver.

## Installation
```shell
go get -u github.com/gopi-frame/database/postgres
```

## Import
```go
import _ "github.com/gopi-frame/database/postgres"
```

## Usage

```go
package main

import (
	"github.com/gopi-frame/database"
	
	_ "github.com/gopi-frame/database/postgres"
)

func main() {
	db, err := database.Connect("postgres", map[string]any{
		"dsn": "user=postgres password=password dbname=postgres sslmode=disable",
    }),
	if err!= nil {
		panic(err)
	}
}
```

## Options

This package uses package [mapstructure](github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [postgres.Config](https://pkg.go.dev/gorm.io/driver/postgres#Config).

### Example
```go
var options = map[string]any{
	"dsn": "user=postgres password=password dbname=postgres sslmode=disable",
	"WithoutQuotingCheck": true,
}
```
