package route

import "net/http"

func Head(rt, handler string) *route {
	return &route{
		method:  http.MethodHead,
		route:   rt,
		handler: handler,
	}
}
