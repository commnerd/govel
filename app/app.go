package app

import (
	"github.com/spf13/viper"
)

type application struct {
	*viper.Viper
}

func New() *application {
	return &application{
		Viper: viper.New(),
	}
}
