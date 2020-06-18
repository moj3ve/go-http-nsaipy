package core

import (
	"fmt"
	_util "go-http/src/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ResolveToJson(uri string) (raw []byte) {
	fmt.Println("[*] $ (DEBUG) Executable located at: " + getExecutableFilepath())
	content, err := ioutil.ReadFile(getExecutableFilepath() + "/contents" + uri + ".json")
	_util.HandleError(err)
	return []byte(content)
}

func getExecutableFilepath() (path string) {
	ex, err := os.Executable()
	_util.HandleError(err)
	exPath := filepath.Dir(ex)
	return exPath
}
