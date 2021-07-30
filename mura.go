package mura

import (
	"reflect"

	"github.com/spf13/viper"
)

// Mura is a stuct to hold viper struct
type Mura struct{ *viper.Viper }

// New is a function to initialize mura struct
func New() *Mura {
	return &Mura{viper.New()}
}

// BindSysEnv is a workaround to make the unmarshal work with environment variables
// reference: https://github.com/spf13/viper/issues/761#issuecomment-626122696
func (m *Mura) BindSysEnv(iface interface{}) {
	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}
	for i := 0; i < ifv.NumField(); i++ {
		field := ifv.Type().Field(i)
		value, ok := field.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		m.BindEnv(value)
	}
}
