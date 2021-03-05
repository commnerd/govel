package Route

import "net/http"

func Head(route, handler string) *route {
	return &route{
		method:  http.MethodHead,
		route:   route,
		handler: handler,
	}
}
