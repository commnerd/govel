package route

import "net/http"

func Head(path, handler string) *routeStruct {
	return craft(path, handler, http.MethodHead)
}
