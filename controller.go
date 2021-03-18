package govel

import (
	"reflect"

	"github.com/commnerd/govel/controller"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/response"
)

type Controller struct {
	*controller.Controller
}

func RegisterController(iController interface{}) {
	tController := reflect.TypeOf(iController)
	vController := reflect.ValueOf(iController)

	_, ok := vController.FieldByName("Controller").Interface().(*Controller)
	if !ok {
		gerror.Throw("Not a controller.")
	}

	ctrl := controller.New()

	for i := 0; i < vController.NumMethod(); i++ {
		method, ok := vController.Method(i).Interface().(func(...interface{}) response.Response)
		if ok {
			ctrl.SetMethod(vController.Method(i).String(), method)
		}
	}

	controller.Set(tController.Name(), ctrl)
}
