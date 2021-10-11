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
	os.Setenv("SERVER_PRODUCTION", "true")
	os.Setenv("PI_CONST", "3.14")
	defer func() {
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("SERVER_PRODUCTION")
		os.Unsetenv("PI_CONST")
	}()

	t.Run("success", func(t *testing.T) {

		type TestENV struct {
			ServerHost       string `env:"SERVER_HOST"`
			ServerPort       int    `env:"SERVER_PORT"`
			ServerProduction bool   `env:"SERVER_PRODUCTION"`

			PiConts float64 `env:"PI_CONST"`
		}

		env := new(TestENV)
		err := Unmarshal(env)
		print(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

	})

	t.Run("success-default", func(t *testing.T) {

		type TestENV struct {
			ServerHost string `env:"SERVER_HOST"`
			DBHost     string `env:"DB_HOST" default:"localhost"`
		}

		env := new(TestENV)
		err := Unmarshal(env)
		print(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

	})

	t.Run("no-env-no-default", func(t *testing.T) {

		type TestENV struct {
			ServerHost string `env:"SERVER_HOST"`
			DBHost     string `env:"DB_HOST" default:"localhost"`
			DBPort     int
		}

		env := new(TestENV)
		err := Unmarshal(env)
		print(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

	})

	t.Run("error-conversion-type", func(t *testing.T) {

		type TestENV struct {
			ServerHost string  `env:"SERVER_HOST"`
			DBHost     string  `env:"DB_HOST" default:"localhost"`
			DBPort     int     `default:"it should be integer"`
			DBSSL      bool    `default:"it should be boolean"`
			PiConts    float64 `default:"it should be float"`
		}

		env := new(TestENV)
		err := Unmarshal(env)
		print(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

	})

	t.Run("error-interface-non-pointer", func(t *testing.T) {

		type TestENV struct{}

		env := TestENV{}
		err := Unmarshal(env)

		if err == nil {
			t.Error("Expectec error")
			t.Fail()
		}

	})

	t.Run("error-unsuported-type", func(t *testing.T) {

		type CustomType struct{}
		type TestENV struct {
			Test CustomType `default:"custom type"`
		}

		env := new(TestENV)
		err := Unmarshal(env)

		if err != nil {
			t.Error("Expectec error")
			t.Fail()
		}

	})
}

func TestMuraWithEnvFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer func() {
			os.Remove(".env")
		}()
		os.WriteFile(".env", []byte("SERVER_PORT=6969"), 0600)
		SetENVPath(".env")

		type TestENV struct {
			ServerHost       string `env:"SERVER_HOST"`
			ServerPort       int    `env:"SERVER_PORT"`
			ServerProduction bool   `env:"SERVER_PRODUCTION"`

			PiConts float64 `env:"PI_CONST"`
		}

		env := new(TestENV)
		err := Unmarshal(env)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

		print(env)
	})
}
