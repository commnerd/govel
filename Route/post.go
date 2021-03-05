package Route

import "net/http"

func Post(route, handler string) *route {
	return &route{
		method:  http.MethodPost,
		route:   route,
		handler: handler,
	}
}
