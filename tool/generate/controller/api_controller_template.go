package controller

const API_CONTROLLER = `package %s

import (
	"Govel"
)

type %s Govel.Controller

func (c *%s) Index() Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Store(request Govel.Request) Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Show(request Govel.Request) Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Update(request Govel.Request) Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Destroy(request Govel.Destroy) Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
