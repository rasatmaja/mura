package mura

import (
	"os"
	"testing"
)

func TestMura(t *testing.T) {
	os.Setenv("SERVER_HOST", "127.0.0.1")

	type test struct {
		Host string `mapstructure:"SERVER_HOST"`
		Port int
	}

	env := new(test)

	mura := New()
	mura.AutomaticEnv()
	mura.BindSysEnv(env)

	mura.ReadInConfig()
	mura.Unmarshal(env)

	if env.Host != os.Getenv("SERVER_HOST") {
		t.Fail()
	}

	// expect port 0 becase this filed doesnt have mapstructure tag
	if env.Port != 0 {
		t.Fail()
	}
}
