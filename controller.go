package govel

import "net/http"

type controllerInterface interface {
	GetRequest() http.Request
}

type Controller struct {
	controllerInterface
}
