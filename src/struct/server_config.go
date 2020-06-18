package _struct

import "net/http"

type ServerConfig struct {
	Addr     string
	Port     uint64
	Handler  http.Handler
	CertFile string
	KeyFile  string
}
