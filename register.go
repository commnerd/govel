package govel

import (
	"github.com/commnerd/govel/app"
	"reflect"
)

func Register(i interface{}) {
	app.Set(reflect.TypeOf(i).Name(), i)
}