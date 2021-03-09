package controller

const WEB_CONTROLLER = `package %s

import (
	"Govel"
)

type %s Govel.Controller

func (c *%s) Index() Govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Create(request Govel.Request) Govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Store(request Govel.Request) Govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Show(request Govel.Request) Govel.Response {
	return response.View("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Edit(request Govel.Request) Govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Update(request Govel.Request) Govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

func (c *%s) Destroy(request Govel.Request) Govel.Response {
	return response.Redirect("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
