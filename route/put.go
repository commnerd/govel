package route

import "net/http"

func Put(path, handler string) *route {
	return craft(path, handler, http.MethodPut)
}
