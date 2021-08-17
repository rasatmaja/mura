<p align="center">
    <img src="https://assets.rasatmaja.com/mura/gopher.webp" width=200> 
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/rasatmaja/mura"><img src="https://godoc.org/github.com/rasatmaja/mura?status.svg"></a>
    <a href="https://github.com/rasatmaja/mura/releases"><img src="https://img.shields.io/github/v/release/rasatmaja/mura"></a>
    <a href=""><img src="https://img.shields.io/github/go-mod/go-version/rasatmaja/mura"></a>
    <a href="https://github.com/rasatmaja/mura/blob/main/LICENSE"><img src="https://img.shields.io/github/license/rasatmaja/mura"></a>
</p>

# 🐍 Mura 
An environment variable reader for Go that bind env value into struct variable, with zero external dependency. The main idea of ​​making this project is the use of `os` and `reflect` packages from Go.

## Overview
This package will provide following features:
- Read system environtment variable bassed on `env` struct tag
- Bind env value into struct 
- Fill struct with default value if env not present

## Installation
To install this package we can run `go` command:
```bash
go get -u github.com/rasatmaja/mura/v2
```

## Usage
```go 
import "github.com/rasatmaja/mura/v2"

type Config struct {
    Host  string `env:"SERVER_HOST" default:"localhost"`
    Port  int    `env:"SERVER_PORT" default:"8080"`
    TLS   bool   `env:"SERVER_TLS" default:"false"`
}

cfg := new(Config)

err := mura.Unmarshal(cfg)
if err != nil {
    panic(err)
}
```

The code above will do:
1. Read environtment variable based on `env` tag defined in struct
2. If the environment variable is found, it will fill the struct value with `env` value
3. If no environment variable found, the struct field will be filled with `default` tag's value 
