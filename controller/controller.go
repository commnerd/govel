package controller

import "github.com/commnerd/govel/response"

type Controller struct {
	methods map[string]func(...interface{}) response.Response
}

func New() *Controller {
	return &Controller{
		methods: make(map[string]func(...interface{}) response.Response),
	}
}

func (c *Controller) SetMethod(label string, method func(...interface{}) response.Response) {
	c.methods[label] = method
}

func (c *Controller) GetMethod(label string) func(...interface{}) response.Response {
	return c.methods[label]
}
