package router

import (
	"github.com/commnerd/govel/app"
	"github.com/commnerd/govel/gerror"
	"github.com/spf13/viper"
)

type Router struct{
	*viper.Viper
}

func New() *Router {
	return &Router{
		Viper: viper.New(),
	}
}

func (r *Router) Register(i interface{}) error {
	return nil
}

var instance *Router

func init() {
	instance = New()
	err := app.Register(instance)
	if err != nil {
		gerror.Check(err)
	}
}