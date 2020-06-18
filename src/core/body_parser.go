package core

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func parseBody(body io.ReadCloser) (rawBody []byte, isJson bool) {
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("[*] isJson?", isJSON(string(raw)))
	//fmt.Println("[*] isJsonString?", isJSONString(string(raw)))

	return raw, isJSON(string(raw))
}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}
