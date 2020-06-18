package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ResolveToJson(uri string) (raw []byte) {
	fmt.Println("[*] $ (DEBUG) Executable located at: " + getExecutableFilepath())

	content, err := ioutil.ReadFile(getExecutableFilepath() + "/contents" + uri + ".json")
	if err != nil {
		fmt.Print("[*] $ (DEBUG) PANIC: ")
		fmt.Println(err)
	}

	return []byte(content)
}

func getExecutableFilepath() (path string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
