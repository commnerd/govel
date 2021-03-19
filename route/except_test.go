package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcept(t *testing.T) {
	rt := craft("/craft/struct/test", "testRoute@Index", "GET")
	assert.IsType(t, &routeStruct{}, rt)
}
