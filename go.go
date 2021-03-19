package govel

import (
	"github.com/commnerd/govel/server"
)

func Go(a app) {
	bootstrap(a)
	server.Serve()
}
