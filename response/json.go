package response

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/commnerd/govel/gerror"
)

func Json(i interface{}) *Response {
	var response *Response

	switch input := i.(type) {
	case string:
		response = &Response{Response: jsonResponseFromString(input)}
	default:
		gerror.Check(errors.New("No handler for passed Json element."))
	}

	return response
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

func jsonResponseFromString(input string) *http.Response {
	return &http.Response{
		Status:           "200 OK",
		StatusCode:       200,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           jsonBuildResponseHeaders(),
		Body:             jsonBuildResponseBody(input),
		ContentLength:    int64(len(input)),
		TransferEncoding: nil,
		Close:            true,
		Uncompressed:     true,
		Trailer:          make(http.Header),
		// Request:          request,
		// TLS:              request.TLS,
	}
}
