package route

import (
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
)

func (r *routeStruct) Name(name string) *routeStruct {
	rtr, ok := app.Get("*router.Router").(*router.Router)
	if !ok {
		gerror.Throw("router not set")
	}
	rtr.RegisterAlias(name, r.path)
	return r
}
