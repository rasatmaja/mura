package mura

import "github.com/spf13/viper"

// Mura is a stuct to hold viper struct
type Mura struct{ *viper.Viper }

// New is a function to initialize mura struct
func New() *Mura {
	return &Mura{viper.New()}
}
