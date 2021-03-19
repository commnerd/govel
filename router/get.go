package router

import "github.com/commnerd/govel/gerror"

func Get() *Router {
	rtr, ok := app.Get("*config.Router").(*Router)

	if !ok {
		gerror.Throw("no registered router")
	}

	return rtr
}
