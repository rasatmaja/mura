package mura

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func init() {
	log.SetPrefix("[ MURA ] ")
}

// Unmarshal is a function to unmarshall value from env to struct value
func Unmarshal(strct interface{}) error {
	iface := reflect.ValueOf(strct)

	// if interface not pointer
	if iface.Kind() != reflect.Ptr {
		// return error
		return fmt.Errorf("interface:%v, isn't pointer", iface.Type().Name())
	}

	iface = iface.Elem()

	for i := 0; i < iface.NumField(); i++ {
		ivalue := iface.Field(i)
		field := iface.Type().Field(i)

		// lookup and get value from env tag in struct field
		env, ok := field.Tag.Lookup("env")

		// if tag env found
		// then bind those env with field struct
		if ok {
			err := bind(ivalue, env)
			// if error is nil then continue into next iteration
			if err == nil {
				continue
			}
			log.Printf("Cannot bind field %s with env:%s, got error: %v", field.Name, env, err)
		}

		// if binding process error then
		// fill struct field with default value
		val, ok := field.Tag.Lookup("default")

		// if default tag found
		if ok {
			// fill struct value with default value
			err := fill(ivalue, val)
			if err != nil {
				log.Printf("Cannot bind field %s with default value (%s), got error: %v", field.Name, val, err)
			}
		}
	}
	return nil
}

var errBindENVNotFound = fmt.Errorf("env not found")

func bind(field reflect.Value, key string) error {
	// lookup env based on param key
	// if env found
	if env, ok := os.LookupEnv(key); ok {
		// then fill field struct with env value
		return fill(field, env)
	}
	// return error not found
	return errBindENVNotFound
}

func fill(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(v)
	case reflect.Int:
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		field.SetInt(v)
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(v)
	default:
		return fmt.Errorf("type:%v, not supported", field.Kind())
	}
	return nil
}
