package route

type application interface {
	Register(interface{}) error
	Get(string) interface{}
}

var app application

func Bind(a interface{}) {
	app = a.(application)
}
