package govel

import (
	"testing"

	"github.com/commnerd/govel/controller"
	"github.com/commnerd/govel/response"
	"github.com/stretchr/testify/assert"
)

type testController struct{ *Controller }

func (c *testController) Index() response.Response {
	return response.Response{}
}

func TestRegisterController(t *testing.T) {
	RegisterController(testController{})

	assert.Equal(t, response.Response{}, controller.Get("testController").GetMethod("Index")())
}
