package route

import "net/http"

func Get(path, handler string) *route {
	return craft(path, handler, http.MethodGet)
}
