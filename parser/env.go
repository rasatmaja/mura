package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var singleton sync.Once

// EnvMaps will hold key and value from .env file
var EnvMaps = map[string]string{}

func init() {
	log.SetPrefix("[ MURA ] ")
}

//ENV parse env file from given path
func ENV(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {

		// Skip line when empty
		v := sc.Text()
		if len(v) != 0 {
			env := strings.Split(v, "=")

			// get env[0] as key and env[1] as value
			EnvMaps[env[0]] = env[1]
		}

	}
	return sc.Err()
}

// GetENV is function to get env maps
func GetENV(path, key string) string {
	if len(path) > 0 {
		singleton.Do(func() {
			if err := ENV(path); err != nil {
				log.Panic(err)
			}
		})
		return EnvMaps[key]
	}
	return ""
}
