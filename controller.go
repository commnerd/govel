package govel

import (
	"reflect"

	"github.com/commnerd/govel/controller"
)

type Controller struct {
	*controller.Controller
}

func RegisterController(iCtrl interface{}) {
	tController := reflect.TypeOf(iCtrl)
	vController := reflect.ValueOf(iCtrl)

	contController := controller.New()

	for i := 0; i < tController.NumMethod(); i++ {
		method := tController.Method(i)
		def := vController.Method(i).Interface()
		contController.SetHandler(method.Name, def)
	}

	controller.Set(tController.Name(), contController)
}
