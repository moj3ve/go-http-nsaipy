package core

import (
	"go-http/src/struct"
	_util "go-http/src/utils"
	"net/http"
	"strconv"
)

func Run(config _struct.ServerConfig) {
	err := http.ListenAndServeTLS(
		config.Addr+":"+strconv.FormatInt(int64(config.Port), 10),
		config.CertFile,
		config.KeyFile,
		config.Handler,
	)
	
	_util.HandlePanic(err)
}
