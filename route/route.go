package route

import "github.com/commnerd/govel/response"

type route struct {
	path      string
	methodMap map[string]func(...interface{}) response.Response
	name      string
	only      []string
}
