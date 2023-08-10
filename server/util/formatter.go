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
	if len(strs) == 0 {
		return ""
	}
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

// GetCompositeKey return a Redis composite key such as "part1:part2:part3:..."
func GetCompositeKey(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}
	var builder strings.Builder
	defer builder.Reset()
	for idx, part := range parts {
		builder.WriteString(part)
		if idx+1 < len(parts) {
			builder.WriteString(":")
		}
	}
	return builder.String()
}

func String2Time(timeStr string) (time.Time, bool) {
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return t, false
	}
	return t, true
}

func String2Int64(str string) (int64, bool) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return i, false
	}
	return i, true
}
