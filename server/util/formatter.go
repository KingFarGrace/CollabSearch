package util

import (
	"strconv"
	"strings"
	"time"
)

// Concat is a function used to concatenate multiple strings.
// Encapsulate the functionality of strings.Builder.
// Allow string, bool and int input.
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

func String2Time(timeStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return t, err
	}
	return t, err
}
