package route

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
)

func craft(path, sControllerHandler, method string) *routeStruct {
	rtr, ok := app.Get("*router.Router").(*router.Router)
	if !ok {
		gerror.Throw("router not set")
	}

	rt := rtr.GetRoute(path)

	aControllerHandler := strings.Split(sControllerHandler, "@")

	sController := aControllerHandler[0]
	sHandler := aControllerHandler[1]

	if rt != nil {
		if rt.GetController() != sController {
			gerror.Throw("controller mismatch for route")
		}

		rs := rt.(*routeStruct)
		rs.methodMap[method] = sHandler
		return rs
	}

	methodMap := make(map[string]string)

	methodMap[method] = sHandler

	rt = &routeStruct{
		path:       path,
		methodMap:  methodMap,
		controller: sController,
		name:       "",
		only:       make([]string, 0),
	}
	rtr.Register(rt)

	rtStruct, ok := rt.(*routeStruct)

	if !ok {
		gerror.Throw(fmt.Sprintf("\"%s\" is not a route.", reflect.TypeOf(rt).Name()))
	}

	return rtStruct
}
