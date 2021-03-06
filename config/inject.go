package config

import (
	"reflect"
)

func Inject(a iApp) {
	app = a
	a.Set(reflect.TypeOf(conf).String(), conf)
}
