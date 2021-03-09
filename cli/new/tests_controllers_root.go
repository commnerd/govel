package new

const TESTS_CONTROLLERS_ROOT = `
package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"bytes"

	"%s/controllers"
)

func TestRootController(t *testing.T) {

	buf := new(bytes.Buffer)
    buf.ReadFrom(controllers.RootController.Index().Body)
    str := buf.String()

	assert.Equal(t, "{\"status\":\"success\",\"data\":\"Wecome!\"}", str)
}
`