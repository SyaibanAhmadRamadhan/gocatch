package gcommon

// Ternary is a function that mimic ternary operator in other language
func Ternary[T any](cond bool, true, false T) T {
	if cond {
		return true
	}

	return false
}
