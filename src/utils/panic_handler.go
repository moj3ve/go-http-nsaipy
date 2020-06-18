package _util

import "fmt"

func HandlePanic(err error) {
	if err != nil {
		fmt.Print("[*] $ (PANIC): ")
		fmt.Println(err)
	}
}
