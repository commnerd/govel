package Route

import "net/http"

func Put(route, handler string) *route {
	return &route{
		method:  http.MethodPut,
		route:   route,
		handler: handler,
	}
}
