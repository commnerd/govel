package router

import (
	"fmt"
	"reflect"

	"github.com/commnerd/govel/gerror"
)

func (r *Router) GetRoute(path string) route {
	iRt := r.Get(path)

	if iRt == nil {
		return nil
	}

	rt, ok := iRt.(route)
	if !ok {
		gerror.Throw(fmt.Sprintf("\"%s\" not a route", reflect.TypeOf(rt).Name()))
	}

	return rt
}
