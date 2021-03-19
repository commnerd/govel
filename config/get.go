package config

import "github.com/commnerd/govel/gerror"

func Get() *Config {
	cfg, ok := app.Get("*config.Config").(*Config)

	if !ok {
		gerror.Throw("no registered config")
	}

	return cfg
}
