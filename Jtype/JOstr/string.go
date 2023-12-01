package JOstr

import (
	"strings"
)

// EmptyString is a constant for empty string
var EmptyString = ""

// LastStringOfSubStr is a function to get last string of substring
func LastStringOfSubStr(str string, substr string) string {
	i := strings.LastIndex(str, substr)
	if i < 0 {
		return str
	}

	return str[i+len(substr):]
}

// FirstCharToLower is a function to convert first character of string to lower case
func FirstCharToLower(str string) string {
	return strings.ToLower(str[:1])
}

// ToLower is a function to convert string to lower case
func ToLower(str string) string {
	str = strings.ToLower(str)
	return str
}
