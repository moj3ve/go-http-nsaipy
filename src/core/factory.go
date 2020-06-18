package core

import (
	_struct "go-http/src/struct"
)

func Start(addr string, port uint64) {
	Run(_struct.ServerConfig{
		Addr:     addr,
		Port:     port,
		Handler:  HttpHandler{},
		CertFile: "cert/server.crt",
		KeyFile:  "cert/server.key",
	})
}

/*
func buildMux() (newMux *http.ServeMux) {
	mux := http.NewServeMux()
	return mux
}

func registerRouteHandlers(mux *http.ServeMux) {
	//mux.HandleFunc("/", HandleServer)
}
*/
