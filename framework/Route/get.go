package Route

import "net/http"

func Get(route, handler string) *route {
	return &route{
		method:  http.MethodGet,
		route:   route,
		handler: handler,
	}
}
