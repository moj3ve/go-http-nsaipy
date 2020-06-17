package core

import (
	"go-http/src/struct"
	"net/http"
	"strconv"
)

func Run(config _struct.ServerConfig) {
	/* TODO: replace nil by real handler */
	http.ListenAndServe(config.Addr+":"+strconv.FormatInt(config.Port, 10), config.Handler)
}
