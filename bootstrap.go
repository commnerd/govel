package govel

import (
	"github.com/commnerd/govel/config"
	"github.com/commnerd/govel/router"
	"github.com/commnerd/govel/server"
)

func bootstrap(a app) {
	config.Bind(a)
	router.Bind(a)
	server.Bind(a)

	config.Get().ReadInConfigs(a.Path("config"))
}
