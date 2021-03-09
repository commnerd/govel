package router

import (
	"net/http"
	"reflect"

	"github.com/commnerd/govel/response"
)

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	response := r.getResponse(request)

	response.Write(w)
}

func (r *Router) getResponse(request *http.Request) response.Response {
	rt := r.GetRoute(request.URL.Path)

	rtHandler := rt.GetHandler(request.Method)

	tHandler := reflect.TypeOf(rtHandler)

	handlerParams := getHandlerParams(tHandler)

	return rtHandler(handlerParams...)
}

func getHandlerParams(rtHandler reflect.Type) []interface{} {
	return []interface{}{rtHandler}
}
