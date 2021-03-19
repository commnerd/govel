package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func New() *Config {
	c := &Config{
		Viper: viper.New(),
	}
	c.SetConfigType("yaml")
	return c
}
