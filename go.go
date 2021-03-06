package govel

import (
	"github.com/commnerd/govel/app"
	"github.com/commnerd/govel/server"
	_ "github.com/commnerd/govel/server"
)

func Go() {
	app.Bootstrap()

	server.Serve()
}
