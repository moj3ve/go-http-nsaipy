package core

import (
	_struct "go-http/src/struct"
	"net/http"
)

func Start(addr string, port int64, handler http.Handler) {
	//mux := buildMux()
	//registerRouteHandlers(mux);
	Run(_struct.ServerConfig{
		Addr:    addr,
		Port:    port,
		Handler: handler,
	})
}

func buildMux() (newMux *http.ServeMux) {
	mux := http.NewServeMux()
	return mux
}

func registerRouteHandlers(mux *http.ServeMux) {
	//mux.HandleFunc("/", HandleServer)
}
