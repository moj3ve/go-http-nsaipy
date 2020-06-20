package _util

import "fmt"

func HandlePanic(err error) (hasError bool){
	if err != nil {
		fmt.Print("[*] $ (PANIC): ")
		fmt.Println(err)
	}

	return err != nil
}
