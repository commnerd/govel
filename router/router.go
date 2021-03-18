package router

import (
	"github.com/spf13/viper"
)

type Router struct {
	*viper.Viper
}

var instance *Router

func init() {
	instance = New()
}
