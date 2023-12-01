package gstr

import (
	"strings"
)

// EmptyString is a constant for empty string
var EmptyString = ""

// GetLastSubstring returns the last substring within a string after the last occurrence of the specified substring.
func GetLastSubstring(str string, substr string) string {
	i := strings.LastIndex(str, substr)
	if i < 0 {
		return str
	}

	return str[i+len(substr):]
}

// LowercaseFirstChar converts the first character of a string to lowercase.
func LowercaseFirstChar(str string) string {
	return strings.ToLower(str[:1])
}

// ToLowercase converts a string to lowercase.
func ToLowercase(str string) string {
	str = strings.ToLower(str)
	return str
}

// ToUppercase converts a string to uppercase.
func ToUppercase(str string) string {
	str = strings.ToUpper(str)
	return str
}

// SplitByWhiteSpace splits a string into substrings separated by white spaces and returns them as a slice.
func SplitByWhiteSpace(s string) []string {
	return strings.Fields(s)
}

// RepeatString repeats a string n times.
func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}

// TrimSpace removes all leading and trailing white spaces from a string.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// CountSubstring counts the number of non-overlapping instances of a substring in a string.
func CountSubstring(s, substr string) int {
	return strings.Count(s, substr)
}
