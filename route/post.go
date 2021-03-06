package route

import "net/http"

func Post(rt, handler string) *route {
	return &route{
		method:  http.MethodPost,
		route:   rt,
		handler: handler,
	}
}
