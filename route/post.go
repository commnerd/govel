package route

import "net/http"

func Post(path, handler string) *routeStruct {
	return craft(path, handler, http.MethodPost)
}
