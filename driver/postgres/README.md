# Postgres
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/database/driver/postgres.svg)](https://pkg.go.dev/github.com/gopi-frame/database/driver/postgres)
[![Test driver postgres](https://github.com/gopi-frame/database/actions/workflows/postgres.yml/badge.svg)](https://github.com/gopi-frame/database/actions/workflows/postgres.yml)
[![codecov](https://codecov.io/gh/gopi-frame/database/graph/badge.svg?token=XRKDX3B3PN&flag=postgres)](https://codecov.io/gh/gopi-frame/database?flags[0]=postgres)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/database/driver/postgres)](https://goreportcard.com/report/github.com/gopi-frame/database/driver/postgres)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

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

This package uses package [mapstructure](https://github.com/go-viper/mapstructure/v2) to parse options.

For more information on the options, see [postgres.Config](https://pkg.go.dev/gorm.io/driver/postgres#Config).

### Example
```go
var options = map[string]any{
	"dsn": "user=postgres password=password dbname=postgres sslmode=disable",
	"WithoutQuotingCheck": true,
}
```
