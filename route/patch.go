package route

import "net/http"

func Patch(path, handler string) *routeStruct {
	return craft(path, handler, http.MethodPatch)
}
