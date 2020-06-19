package main

import (
	"flag"
	"fmt"
	"go-http/src/core"
)

func main() {
	printHeader()
	port := flag.Uint64("p", 443, "port")
	flag.Parse()
	core.Start("", *port)
	return
}

func printHeader() {
	fmt.Println("[*] go-http-nsaipy")
	fmt.Println("[*] http utility written in golang - not sure about its purpose yet")
	fmt.Println("[*] (yes the name is crap but I will care later)")
	fmt.Println("[*] Web:    https://danschmit.dev")
	fmt.Println("[*] GitHub: https://github.com/woodfairy")
	fmt.Println("[*] Copyright (c) 2020 by woodfairy")
	fmt.Println()
}
