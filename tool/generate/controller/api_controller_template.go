package controller

const API_CONTROLLER = `package %s

import (
	"govel"
)

type %s govel.Controller

func (c *%s) Index() *govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Store(request govel.Request) *govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Show(request govel.Request) *govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Update(request govel.Request) *govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Destroy(request govel.Destroy) *govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
