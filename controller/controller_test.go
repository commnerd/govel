package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Equal(t, &Controller{methods: make(map[string]interface{})}, New())
}

func testMethod() string { return "tested" }

func TestSetHandler(t *testing.T) {
	ctrl := New()

	ctrl.SetHandler("test", testMethod)

	iTestMethod := ctrl.methods["test"]
	method, ok := iTestMethod.(func() string)
	if !ok {
		panic("not a method.")
	}

	expected := testMethod()
	actual := method()

	assert.Equal(t, expected, actual)
}

func TestGetHandler(t *testing.T) {
	ctrl := New()

	ctrl.methods["test"] = testMethod

	iTestMethod := ctrl.GetHandler("test")
	method, ok := iTestMethod.(func() string)
	if !ok {
		panic("not a method.")
	}

	expected := testMethod()
	actual := method()

	assert.Equal(t, expected, actual)
}
