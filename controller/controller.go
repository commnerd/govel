package controller

type Controller struct {
	methods map[string]interface{}
}

func New() *Controller {
	return &Controller{
		methods: make(map[string]interface{}),
	}
}

func (c *Controller) SetHandler(label string, method interface{}) {
	c.methods[label] = method
}

func (c *Controller) GetHandler(label string) interface{} {
	return c.methods[label]
}
