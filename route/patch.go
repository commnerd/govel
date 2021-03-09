package route

import "net/http"

func Patch(path, handler string) *route {
	return craft(path, handler, http.MethodPatch)
}
