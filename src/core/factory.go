package core

import (
	_struct "go-http/src/struct"
	"net/http"
)

func Start(addr string, port int64, handler http.Handler) {
	Run(_struct.ServerConfig{
		Addr:    addr,
		Port:    port,
		Handler: handler,
	})
}
