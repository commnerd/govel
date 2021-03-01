package Route

import "net/http"

func Patch(route, handler string) *route {
	return &route{
		method:  http.MethodPatch,
		route:   route,
		handler: handler,
	}
}
