package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type sampleStruct struct{}

func TestGetStructName(t *testing.T) {
	assert.Equal(t, "sampleStruct", GetStructName(interface{}(sampleStruct{})))
}