package core

import (
	_util "go-http/src/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Resolve(uri string) (raw []byte) {
	//fmt.Println("[*] $ (DEBUG) Executable located at: " + getExecutableFilepath())
	content, err := ioutil.ReadFile(getExecutableFilepath() + "/hosts" + uri)
	hasError := _util.HandlePanic(err)
	if(hasError) {
		fallback, fallbackErr := ioutil.ReadFile(getExecutableFilepath() + "/hosts/fallback.json")
		_util.HandlePanic(fallbackErr)
		return fallback
	}
	return content
}

func getExecutableFilepath() (path string) {
	ex, err := os.Executable()
	_util.HandlePanic(err)
	exPath := filepath.Dir(ex)
	return exPath
}
