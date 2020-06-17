package core

import (
	"fmt"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[*] $ Incoming " + request.Method + " Request")
	fmt.Println("[*] $ User Agent: " + request.Header.Get("User-Agent"))
	data := []byte("<h1>go-http-nsaipy!</h1><p>It works!</p><a target=\"_blank\" href=\"https://github.com/woodfairy\">GitHub</a> - <a target=\"_blank\" href=\"https://danschmit.dev\">Web / Matrix</a>")
	writer.Write(data)
	return
}
