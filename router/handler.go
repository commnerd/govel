package router

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/commnerd/govel/controller"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/response"
)

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	response := r.getResponse(request)

	response.Write(w)
}

func (r *Router) getResponse(request *http.Request) response.Response {
	rt := r.GetRoute(request.URL.Path)

	sController := rt.GetController()

	ctrl := controller.Get(sController)

	sHandler := rt.GetHandler(request.Method)

	tCtrl := reflect.TypeOf(ctrl)
	tHandler, ok := tCtrl.MethodByName(sHandler)
	if !ok {
		gerror.Throw(fmt.Sprintf("\"%s\" is not a handler on \"%s\"", sHandler, tCtrl.Name()))
	}

	handlerParams := getHandlerParams(tHandler, request)

	resp := tHandler.Func.Call(handlerParams)

	return responseFromValues(resp)
}

func getHandlerParams(mHandler reflect.Method, req *http.Request) []reflect.Value {
	params := make([]reflect.Value, 0)

	for i := 0; i < mHandler.Type.NumIn(); i++ {
		switch mHandler.Type.In(i).Name() {
		case "*http.Request":
			params = append(params, reflect.ValueOf(req))
		case "request.Request":
			params = append(params, reflect.ValueOf(req))
		default:
			param := app.Get(mHandler.Type.In(i).Name())
			params = append(params, reflect.ValueOf(param))
		}
	}
	return params
}

func responseFromValues(responseValues []reflect.Value) response.Response {
	if len(responseValues) != 1 {
		gerror.Throw("improper implementation of handler")
	}

	resp, ok := responseValues[0].Interface().(response.Response)

	if !ok {
		gerror.Throw("improper implementation of handler")
	}

	return resp
}
