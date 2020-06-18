package core

import (
	"go-http/src/struct"
	_util "go-http/src/utils"
	"net/http"
	"strconv"
)

func Run(config _struct.ServerConfig) {
	/* TODO: replace nil by real handler */
	err := http.ListenAndServeTLS(config.Addr+":"+strconv.FormatInt(config.Port, 10), "server.crt", "server.key", config.Handler)
	_util.HandlePanic(err)
}
