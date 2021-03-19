package route

import "net/http"

func Get(path, handler string) *routeStruct {
	return craft(path, handler, http.MethodGet)
}
