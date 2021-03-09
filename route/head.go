package route

import "net/http"

func Head(path, handler string) *route {
	return craft(path, handler, http.MethodHead)
}