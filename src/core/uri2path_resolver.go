package core

import (
	_util "go-http/src/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ResolveToJson(uri string) (raw []byte) {
	//fmt.Println("[*] $ (DEBUG) Executable located at: " + getExecutableFilepath())
	content, err := ioutil.ReadFile(getExecutableFilepath() + "/hosts" + uri + ".json")
	_util.HandlePanic(err)
	return content
}

func getExecutableFilepath() (path string) {
	ex, err := os.Executable()
	_util.HandlePanic(err)
	exPath := filepath.Dir(ex)
	return exPath
}
