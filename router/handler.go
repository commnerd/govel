package router

import (
	"github.com/commnerd/govel/gerror"
	"net/http"
	"reflect"
	"fmt"
)

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
		response := r.getResponse(request)

		response.Write(w)
}

func (r *Router) getResponse(request *http.Request) *http.Response {
	rt, h := r.mapRoute(request.URL.Path)
}

func (r *Router) mapRoute(path string) handler {
	iRoute := r.Get(path)

	rt, ok := iRoute.(route)
	if !ok {
		gerror.Throw(fmt.Sprintf("Improperly formed route: %s.", reflect.TypeOf(iRoute).Name()))
	}

	handlerLabel := rt.GetHandler()
}