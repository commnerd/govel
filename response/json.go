package response

import (
	"github.com/commnerd/govel/gerror"
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
	"io"
)

func Json(i interface{}) *Response {
	switch input := i.(type) {
	case string:
		return &Response{Response: jsonResponseFromString(input)}
	default:
		gerror.Check(errors.New("No handler for passed Json element."))
	}
}

func jsonBuildResponseHeaders(i ...interface{}) http.Header {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	return headers
}

func jsonBuildResponseBody(i interface{}) io.ReadCloser {
	switch input := i.(type) {
	case string:
		return ioutil.NopCloser(bytes.NewReader([]byte(input)))
	default:
		panic("JSON cannot be parsed.")
	}
}

func jsonResponseFromString(input string, request *http.Request) *http.Response {
	return &http.Response{
		Status: "200 OK",
    	StatusCode: 200,
		Proto: "HTTP/1.0",
    	ProtoMajor: 1,
    	ProtoMinor: 0,
		Header: jsonBuildResponseHeaders(),
    	Body: jsonBuildResponseBody(input),
		ContentLength: len(input),
		TransferEncoding: nil,
		Close: true,
		Uncompressed: true,
		Trailer: make(http.Header),
		Request: request,
		TLS: request.TLS,
	}
}