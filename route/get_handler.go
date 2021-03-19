package route

import (
	"strings"
)

func (rt *routeStruct) GetHandler(sHandler string) string {
	structMethod := strings.Split(sHandler, "@")

	return rt.methodMap[structMethod[1]]
}
