package mura

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func print(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println("============= DEBUG =============")
	fmt.Printf("%s \n", s)
	fmt.Println("=================================")
}

func TestMura(t *testing.T) {
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "8080")

	t.Run("success", func(t *testing.T) {

		type TestENV struct {
			ServerHost string `env:"SERVER_HOST"`
			ServerPort int    `env:"SERVER_PORT"`
		}

		env := new(TestENV)
		err := Unmarshal(env)
		print(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

	})
}
