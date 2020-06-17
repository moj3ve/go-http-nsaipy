package _struct

import (
	"io"
	"net/http"
	"time"
)

type RequestWrapper struct {
	Method  string
	Target  string
	Body    io.Reader
	Timeout uint64
	Headers map[string]string
}

func Build(wrapper RequestWrapper) (httpRequest http.Request, httpClient http.Client) {
	client := buildClient(time.Duration(wrapper.Timeout))
	request, _ := http.NewRequest(wrapper.Method, wrapper.Target, wrapper.Body)
	for key, value := range wrapper.Headers {
		request.Header.Add(key, value)
	}

	return *request, client
}

func buildClient(timeout time.Duration) (client http.Client) {
	return http.Client{
		Timeout: timeout * time.Second,
	}
}
