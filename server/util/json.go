package util

import (
	"encoding/json"
	"os"
)

// ParseJSONFile can parse a .json file and inject the data into a struct.
func ParseJSONFile(path string, obj interface{}) {
	data, err := os.ReadFile(path)
	if nil != err {
		PlainErrorLogger(err, "ParseJSONFile()")
		return
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		PlainErrorLogger(err, "json.Unmarshal()")
		return
	}
}
