package parser

import (
	"os"
	"sync"
	"testing"
)

func TestENVError(t *testing.T) {
	t.Run("error-env-not-found", func(t *testing.T) {
		singleton = sync.Once{}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code didnt panic")
				t.Fail()
			}
		}()
		GetENV(".env", "SERVER_PORT")
	})

	t.Run("empty-path", func(t *testing.T) {
		singleton = sync.Once{}
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code did panic")
				t.Fail()
			}
		}()

		env := GetENV("", "SERVER_PORT")

		if len(env) != 0 {
			t.Error("env SERVER_PORT should be empty")
			t.Fail()
		}
	})

	t.Run("success", func(t *testing.T) {
		singleton = sync.Once{}
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code did panic")
				t.Fail()
			}

			os.Remove(".env")
		}()
		os.WriteFile(".env", []byte("SERVER_PORT=123"), 0600)

		env := GetENV(".env", "SERVER_PORT")

		if env != "123" {
			t.Error("env SERVER_PORT should be 123")
			t.Fail()
		}
	})

	t.Run("success-split-string", func(t *testing.T) {
		singleton = sync.Once{}
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("The code did panic")
				t.Fail()
			}

			os.Remove(".env")
		}()
		os.WriteFile(".env", []byte("SERVER_URL=http://hostname.com?page=2&limit=10"), 0600)

		env := GetENV(".env", "SERVER_URL")

		if env != "http://hostname.com?page=2&limit=10" {
			t.Error("env SERVER_URL should be http://hostname.com?page=2&limit=10")
			t.Fail()
		}
	})
}
