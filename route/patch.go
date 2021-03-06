package route

import "net/http"

func Patch(rt, handler string) *route {
	return &route{
		method:  http.MethodPatch,
		route:   rt,
		handler: handler,
	}
}
