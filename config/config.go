package config

import (
	"github.com/commnerd/govel/app"
	"github.com/commnerd/govel/gerror"
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

	err := app.Register(instance)
	gerror.Check(err)

	build()
}
