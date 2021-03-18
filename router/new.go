package router

import "github.com/spf13/viper"

func New() *Router {
	return &Router{
		Viper: viper.New(),
	}
}
