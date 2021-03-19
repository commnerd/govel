package route

import (
	"net/url"

	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
)

type routeStruct struct {
	path       string
	methodMap  map[string]string
	controller string
	name       string
	only       []string
}

func Route(name string) *url.URL {
	rtr, ok := app.Get("*router.Router").(*router.Router)
	if !ok {
		gerror.Throw("router not set")
	}
	rt := rtr.GetRoute(name)
	return routeToURL(rt)
}

func Path(path string) *url.URL {
	rtr, ok := app.Get("*router.Router").(*router.Router)
	if !ok {
		gerror.Throw("router not set")
	}
	rt := rtr.GetRoute(path)
	return routeToURL(rt)
}

func routeToURL(rt interface{ GetPath() string }) *url.URL {
	url, err := url.Parse(rt.GetPath())
	gerror.Check(err)
	return url
}
