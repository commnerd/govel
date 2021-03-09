package route

import (
	"strings"

	"github.com/commnerd/govel/response"
)

func (rt *route) GetHandler(sHandler string) func(...interface{}) response.Response {
	structMethod := strings.Split(sHandler, "@")

	return rt.methodMap[structMethod[1]]
}
