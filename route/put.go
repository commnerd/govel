package route

import "net/http"

func Put(path, handler string) *routeStruct {
	return craft(path, handler, http.MethodPut)
}
