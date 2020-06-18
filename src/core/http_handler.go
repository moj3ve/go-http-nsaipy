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
	fmt.Println("[*] $ User-Agent: " + request.Header.Get("User-Agent"))
	fmt.Println("[*] $ Content-Type: " + request.Header.Get("Content-Type"))

	//data := []byte("<h1>go-http-nsaipy!</h1><p>It works!</p><a target=\"_blank\" href=\"https://github.com/woodfairy\">GitHub</a> - <a target=\"_blank\" href=\"https://danschmit.dev\">Web / Matrix</a>")

	header := writer.Header()
	header.Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	rawBody, isJson := parseBody(request.Body)
	if(isJson) {
		//TODO: parse JSON
		fmt.Println("[*] $ JSON Body type recognized")
		fmt.Println("[*] $ content: " + string(rawBody))

	} else {
		fmt.Println("[*] $ No JSON type detected. Interpreting as string.")
		fmt.Println("[*] $ String parsed: " + string(rawBody))
	}

	fmt.Println()
	writer.Write(rawBody)
	return
}
