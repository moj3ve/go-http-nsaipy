package main

import (
	"fmt"
	"go-http/src/core"
	"os"
	"strconv"
)

func main() {
	printHeader()
	args := os.Args
	var port uint64

	if(len(args) > 1) {
		port, _ = strconv.ParseUint(args[1], 10, 64)
	} else {
		port = 443
	}

	core.Start("", port)
	return;
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
