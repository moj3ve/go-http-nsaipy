package core

import (
	"encoding/json"
	_util "go-http/src/utils"
	"io"
	"io/ioutil"
)

func parseBody(body io.ReadCloser) (rawBody []byte, isJson bool) {
	raw, err := ioutil.ReadAll(body)
	_util.HandlePanic(err)
	return raw, isJSON(string(raw))
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
