package util

import (
	"github.com/bytedance/sonic"
	"os"
)

// ParseJSONFile can parse a .json file and inject the data into a struct.
func ParseJSONFile(path string, obj interface{}) {
	data, err := os.ReadFile(path)
	if nil != err {
		ErrorLogger(err, "ParseJSONFile()")
		return
	}
	err = sonic.Unmarshal(data, obj)
	if err != nil {
		ErrorLogger(err, "sonic.Unmarshal()")
		return
	}
}
