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