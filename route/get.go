package route

import "net/http"

func Get(rt, handler string) *route {
	return &route{
		method:  http.MethodGet,
		route:   rt,
		handler: handler,
	}
}
