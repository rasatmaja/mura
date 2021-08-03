package mura

import (
	"reflect"
	"strconv"
)

// FillDefault is a method to fill struct value from its own default tag
func (m *Mura) FillDefault(iface interface{}) error {
	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}
	for i := 0; i < ifv.NumField(); i++ {
		v := ifv.Field(i)
		t := ifv.Type().Field(i)
		tv, ok := t.Tag.Lookup("default")
		if !ok {
			continue
		}

		switch v.Kind() {
		case reflect.String:
			ifv.Field(i).SetString(tv)
		case reflect.Bool:
			v, err := strconv.ParseBool(tv)
			if err != nil {
				continue
			}
			ifv.Field(i).SetBool(v)
		case reflect.Int:
			v, err := strconv.ParseInt(tv, 10, 32)
			if err != nil {
				continue
			}
			ifv.Field(i).SetInt(v)
		default:
			continue
		}
	}
	return nil
}
