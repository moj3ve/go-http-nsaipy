package _util

import "fmt"

func HandlePanic(err error) {
	if err != nil {
		fmt.Print("[*] $ (DEBUG) PANIC: ")
		fmt.Println(err)
	}
}
