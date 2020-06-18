package _util

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Print("[*] $ (DEBUG) PANIC: ")
		fmt.Println(err)
	}
}
