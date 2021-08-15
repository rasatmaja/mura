package mura

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Unmarshal is a function to unmarshall value from env to struct value
func Unmarshal(strct interface{}) error {
	iface := reflect.ValueOf(strct)

	// Checking interface kind,
	// if iterface is pointer
	// then get interface element
	if iface.Kind() == reflect.Ptr {
		iface = iface.Elem()
	}

	for i := 0; i < iface.NumField(); i++ {
		field := iface.Field(i)

		// lookup and get value from env tag in struct field
		tag, ok := iface.Type().Field(i).Tag.Lookup("env")

		// if tag env not found continiue iteration
		// into next field
		if !ok {
			continue
		}

		bind(field, tag)
	}
	return nil
}

var errBindENVNotFound = fmt.Errorf("env not found")

func bind(field reflect.Value, key string) error {
	env, ok := os.LookupEnv(key)
	if !ok {
		return errBindENVNotFound
	}
	return fillField(field, env)
}

func fillField(field reflect.Value, value string) error {
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
	}

	return nil
}
