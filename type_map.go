package govel

import (
	"github.com/commnerd/govel/request"
	"github.com/commnerd/govel/response"
)

type Request struct{ *request.Request }
type Response struct{ *response.Response }
