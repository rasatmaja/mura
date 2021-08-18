package main

import (
	"encoding/json"
	"fmt"

	"github.com/rasatmaja/mura/v2"
)

// Config ...
type Config struct {
	Host string `env:"SERVER_HOST" default:"localhost"`
	Port int    `env:"SERVER_PORT" default:"8080"`
	TLS  bool   `env:"SERVER_TLS" default:"true"`
}

func main() {
	cfg := new(Config)

	err := mura.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	print(cfg)
}

func print(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println("============== ENV ==============")
	fmt.Printf("%s \n", s)
	fmt.Println("=================================")
}
