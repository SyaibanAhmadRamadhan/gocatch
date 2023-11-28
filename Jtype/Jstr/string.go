package Jstr

import (
	"strings"
)

func LastStringOfSubStr(str string, substr string) string {
	i := strings.LastIndex(str, substr)
	if i < 0 {
		return str
	}

	return str[i+len(substr):]
}

func FirstCharToLower(str string) string {
	return strings.ToLower(str[:1])
}

func ToLower(str string) string {
	str = strings.ToLower(str)
	return str
}
