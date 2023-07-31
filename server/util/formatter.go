package util

import (
	"strconv"
	"strings"
)

// Concat is a function used to concatenate multiple strings.
// Encapsulate the functionality of strings.Builder.
func Concat(strs ...interface{}) string {
	var builder strings.Builder
	defer builder.Reset()
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		if value, ok := str.(string); ok {
			builder.WriteString(value)
			continue
		}
		if value, ok := str.(bool); ok {
			if value {
				builder.WriteString("true")
			} else {
				builder.WriteString("false")
			}
			continue
		}
		if value, ok := str.(int); ok {
			builder.WriteString(strconv.Itoa(value))
			continue
		}
	}
	return builder.String()
}
