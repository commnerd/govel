package router

import "github.com/commnerd/govel/gerror"

func (r *Router) GetRoute(path string) route {
	rt, ok := r.Viper.Get(path).(route)

	if !ok {
		gerror.Throw("Not a route.")
	}

	return rt
}

var GetRoute = instance.GetRoute
