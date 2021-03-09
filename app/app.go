package app

import (
	"github.com/spf13/viper"
)

type application struct{
	*viper.Viper
}

func New() *application {
	return &application{
		Viper: viper.New(),
	}
}

func (a *application) Register(i interface{}) error {
	return nil
}

var Register = instance.Register

var instance *application

func init() {
	instance = New()
}

var Get = application.Get
var Set = application.Set
