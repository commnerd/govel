package route

import "github.com/commnerd/govel/router"

func craft(path, handler, method string) *route {
	rt := &route{
		path:    path,
		method: method,
		handler: handler,
	}

	router.Register(rt)

	return rt
}