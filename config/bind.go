package config

type application interface {
	Register(interface{}) error
	Get(string) interface{}
}

var app application

func Bind(a interface{}) {
	app = a.(application)

	c := New()

	app.Register(c)
}
