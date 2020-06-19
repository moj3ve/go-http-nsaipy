package core

import (
	"fmt"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	processRequest(writer, request)
	writer.Write(buildResponse(request))
	fmt.Println("-----------------------------------------------------------")
	return
}

func processRequest(writer http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)
	setResponseHeaders(writer)
}

func buildResponse(request *http.Request) (raw []byte) {
	if request.RequestURI == "/" {
		request.RequestURI = "/index"
	}
	return Resolve("/" + request.Host + request.RequestURI + ".json")
}

func printRequestInfo(request *http.Request) {
	fmt.Println("[*] $ Incoming " + request.Method + " Request")
	fmt.Println("[*] $ Host: " + request.Host)
	fmt.Println("[*] $ Remote Addr: " + request.RemoteAddr)
	fmt.Println("[*] $ Request URI: " + request.RequestURI)
	fmt.Println("[*] $ User-Agent: " + request.Header.Get("User-Agent"))
	fmt.Println("[*] $ Content-Type: " + request.Header.Get("Content-Type"))
	rawBody, isJson := parseBody(request.Body)
	if isJson {
		fmt.Println("[*] $ JSON Body type recognized")
		fmt.Println("[*] $ Request Body: \n" + string(rawBody))
	} else {
		fmt.Println("[*] $ No JSON type detected. Interpreting as string.")
		fmt.Println("[*] $ Request Body: \n" + string(rawBody))
	}
}

func setResponseHeaders(writer http.ResponseWriter) {
	header := writer.Header()
	header.Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
