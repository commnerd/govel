package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var instance *Config

func New() *Config {
	return &Config{
		Viper: viper.New(),
	}
}

func init() {
	instance = New()
	instance.SetConfigType("yaml")
}
