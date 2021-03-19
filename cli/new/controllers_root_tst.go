package new

const CONTROLLERS_ROOT_TEST = `package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"bytes"
)

func TestRootController(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(RootController{}.Index().Body)
	str := buf.String()

	assert.Equal(t, "{\"status\":\"success\",\"data\":\"Wecome!\"}", str)
}
`
