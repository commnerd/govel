package controller

import (
	"github.com/commnerd/govel/gerror"
	"github.com/spf13/viper"
)

type Registry struct{ *viper.Viper }

var instance *Registry

func init() {
	instance = NewRegistry()
}

func NewRegistry() *Registry {
	return &Registry{
		Viper: viper.New(),
	}
}

func Set(structLabel string, c *Controller) {
	instance.Set(structLabel, c)
}

func Get(structLabel string) *Controller {
	ctrl, ok := instance.Get(structLabel).(*Controller)
	if !ok {
		gerror.Throw("Something went wrong trying to pull the controller.")
	}
	return ctrl
}
