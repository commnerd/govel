package server

import "github.com/commnerd/govel/response"

type serverRoute interface {
	GetPath() string
	GetHandler(string) func(...interface{}) response.Response
}
