package route

import (
	"strings"

	"github.com/commnerd/govel/response"
	"github.com/commnerd/govel/router"
)

func craft(path, sHandler, method string) *route {
	rt := &route{
		path:      path,
		methodMap: map[string]func(...interface{}) response.Response{method: mapStringToHandler(path, sHandler)},
		name:      "",
		only:      make([]string, 0),
	}

	router.Register(rt)

	return rt
}

func mapStringToHandler(path, sHandler string) func(...interface{}) response.Response {
	structMethod := strings.Split(sHandler, "@")

	rt := router.Get(path).(route)

	return rt.methodMap[structMethod[1]]
}
