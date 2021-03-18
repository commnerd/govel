package controller

const WEB_CONTROLLER = `package %s

import (
	"govel"
)

type %s govel.Controller

func (c *%s) Index() *govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Create(request govel.Request) *govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Store(request govel.Request) *govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Show(request govel.Request) *govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Edit(request govel.Request) *govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Update(request govel.Request) *govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Destroy(request govel.Request) *govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
