package router

import (
	"github.com/commnerd/govel/gerror"
)

type application interface {
	Register(interface{}) error
	Get(string) interface{}
}

var app application

func Bind(a interface{}) {
	app = a.(application)

	err := app.Register(New())

	gerror.Check(err)
}
