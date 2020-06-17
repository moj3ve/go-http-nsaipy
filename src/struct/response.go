package _struct

import (
	"io"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Proto      string
	ProtoMajor int
	ProtoMinor int
	Header     http.Header
	Body       io.ReadCloser
	Request    *http.Request
}
