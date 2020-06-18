package core

import (
	"fmt"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[*] $ Incoming " + request.Method + " Request")
	//fmt.Println("[*] $ Host: " + request.Host)
	//fmt.Println("[*] $ Remote Addr: " + request.RemoteAddr)
	fmt.Println("[*] $ Request URI: " + request.RequestURI)
	fmt.Println("[*] $ User Agent: " + request.Header.Get("User-Agent"))
	fmt.Println()
	data := []byte("<h1>go-http-nsaipy!</h1><p>It works!</p><a target=\"_blank\" href=\"https://github.com/woodfairy\">GitHub</a> - <a target=\"_blank\" href=\"https://danschmit.dev\">Web / Matrix</a>")
	writer.Write(data)
	return
}
