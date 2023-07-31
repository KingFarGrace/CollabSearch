package util

import "strings"

// Concat is a function used to concatenate multiple strings.
// Encapsulate the functionality of strings.Builder.
func Concat(strs ...string) string {
	var builder strings.Builder
	defer builder.Reset()
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}
