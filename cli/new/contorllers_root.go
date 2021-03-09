package new

const CONTROLLERS_ROOT = `package controllers

import (
	"govel"
	"govel/response"
)

type RootController govel.Controller

func (c *RootController) Index() govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}
`
