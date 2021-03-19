package route

import (
	"os"
	"testing"

	gvlApp "github.com/commnerd/govel/app"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
	"github.com/stretchr/testify/assert"
)

type testRoute routeStruct

func (r testRoute) GetController() string {
	return "testController"
}

func (r testRoute) GetPath() string {
	return "/some/test/path"
}

func (r testRoute) GetHandler(sHandler string) string {
	return "testHandler"
}

func TestMain(m *testing.M) {
	app = gvlApp.New()
	router.Bind(app)

	rtr, ok := app.Get("*router.Router").(*router.Router)
	if !ok {
		gerror.Throw("router not set")
	}

	err := rtr.Register(testRoute{})
	gerror.Check(err)

	os.Exit(m.Run())
}

func TestCraftReturnStruct(t *testing.T) {
	rt := craft("/craft/struct/test", "testRoute@Index", "GET")
	assert.IsType(t, &routeStruct{}, rt)
	assert.Equal(t, 1, len(rt.methodMap))

}

func TestCraftSecondMethod(t *testing.T) {
	rt := craft("/craft/second/method/test", "testRoute@Index", "GET")
	craft("/craft/second/method/test", "testRoute@Delete", "DELETE")

	assert.Equal(t, 2, len(rt.methodMap))
}
