package _struct

import "net/http"

type ServerConfig struct {
	Addr    string
	Port    int64
	Handler http.Handler
}
