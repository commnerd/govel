package govel

import (
	"reflect"
	"testing"

	"github.com/commnerd/govel/controller"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/response"
	"github.com/stretchr/testify/assert"
)

type testController Controller

func (c testController) Index() response.Response {
	return response.Response{}
}

func TestRegisterController(t *testing.T) {
	ctrl := testController{}

	RegisterController(ctrl)

	expected := ctrl.Index()
	iMethod := controller.Get("testController").GetHandler("Index")
	actual, ok := reflect.ValueOf(iMethod).Call([]reflect.Value{})[0].Interface().(response.Response)
	if !ok {
		gerror.Throw("Not a proper call to method")
	}
	assert.Equal(t, expected, actual)
}
