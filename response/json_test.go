package response

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonResponseFromString(t *testing.T) {
	resp := jsonResponseFromString("{\"test\":\"out\"}")

	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "HTTP/1.0", resp.Proto)
	assert.Equal(t, 1, resp.ProtoMajor)
	assert.Equal(t, 0, resp.ProtoMinor)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "{\"test\":\"out\"}", string(b))
	assert.Equal(t, int64(14), resp.ContentLength)
	assert.Equal(t, []string([]string(nil)), resp.TransferEncoding)
	assert.Equal(t, true, resp.Close)
	assert.Equal(t, true, resp.Uncompressed)
	assert.Equal(t, http.Header{}, resp.Trailer)
	assert.Equal(t, (*http.Request)(nil), resp.Request)
	assert.Equal(t, (*tls.ConnectionState)(nil), resp.TLS)
}

func TestJsonBuildResponseBody(t *testing.T) {
	rc := jsonBuildResponseBody("{\"foo\":\"bar\"}")

	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	assert.Equal(t, "{\"foo\":\"bar\"}", buf.String())

}

func TestJsonBuildResponseHeaders(t *testing.T) {
	header := jsonBuildResponseHeaders()

	assert.Equal(t, "application/json", header.Get("Content-Type"))
}

func TestJson(t *testing.T) {
	jsonResp := Json("{\"test\":\"out\"}")
	assert.IsType(t, Response{}, jsonResp)

	resp := jsonResp.Response

	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "HTTP/1.0", resp.Proto)
	assert.Equal(t, 1, resp.ProtoMajor)
	assert.Equal(t, 0, resp.ProtoMinor)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "{\"test\":\"out\"}", string(b))
	assert.Equal(t, int64(14), resp.ContentLength)
	assert.Equal(t, []string([]string(nil)), resp.TransferEncoding)
	assert.Equal(t, true, resp.Close)
	assert.Equal(t, true, resp.Uncompressed)
	assert.Equal(t, http.Header{}, resp.Trailer)
	assert.Equal(t, (*http.Request)(nil), resp.Request)
	assert.Equal(t, (*tls.ConnectionState)(nil), resp.TLS)
}
