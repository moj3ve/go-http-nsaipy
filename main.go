package main

import (
	"fmt"
	"go-http/src/core"
)

func main() {
	printHeader()
	core.Start("", 8080, core.HttpHandler{})
	return;
}

func printHeader() {
	fmt.Println("[*] H-olla")
	fmt.Println("[*] T-he Woodfairy")
	fmt.Println("[*] T")
	fmt.Println("[*] P")
	fmt.Println("[*] http utility written in golang by woodfairy")
	fmt.Println("[*] (yes the name is crap but I will care later)")
	fmt.Println("[*] Web:    https://danschmit.dev")
	fmt.Println("[*] GitHub: https://github.com/woodfairy")
	fmt.Println("[*] Copyright (c) 2020 by woodfairy")
	fmt.Println()
}
