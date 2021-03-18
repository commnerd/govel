package router

import "github.com/commnerd/govel/response"

type route interface {
	GetPath() string
	GetHandler(string) func(...interface{}) response.Response
}
