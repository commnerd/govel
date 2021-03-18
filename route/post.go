package route

import "net/http"

func Post(path, handler string) *route {
	return craft(path, handler, http.MethodPost)
}