package new

const CONTROLLERS_ROOT = `
package controllers

import (
	"Govel"
	"Govel/response"
)

type RootController Govel.Controller

func (c *RootController) Index() Govel.Response {
	return response.Json("{\"status\":\"success\",\"data\":\"Wecome!\"}")
}

`
