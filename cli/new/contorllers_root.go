package new

const CONTROLLERS_ROOT = `package controllers

import (
	"govel"
	"govel/response"
)

type RootController struct{ govel.Controller }

func (c *RootController) Index() response.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
