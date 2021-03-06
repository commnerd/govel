package route

import "net/http"

func Put(rt, handler string) *route {
	return &route{
		method:  http.MethodPut,
		route:   rt,
		handler: handler,
	}
}
