package JOmap

import (
	"strconv"
	"strings"
)

// SA is a type alias for a map of string keys to any values.
type SA map[string]any

// JoinKey is a method receiver for SA type.
// It accepts a variable number of string parameters.
// It concatenates the keys of the map, prefixing each key with the first parameter (`ps[0]` or `prefix`)
// and separating each key with the second parameter (`ps[1]` or `separator`).
// If the map itself is nil, it will return an empty string.
// If the prefix is not supplied, it will default to an empty string.
// If the separator is not defined, it will use comma followed by a space.
func (s SA) JoinKey(ps ...string) string {
	if s == nil {
		return ""
	}

	prefix := ""
	separator := ", "

	if ps != nil {
		if len(ps) >= 1 {
			prefix = ps[0]
		}

		if len(ps) >= 2 {
			separator = ps[1]
		}
	}

	var saArr []string
	for k, _ := range s {
		saArr = append(saArr, prefix+k)
	}

	return strings.Join(saArr, separator)
}

func (s SA) Merge(maps ...map[string]any) {
	for i, m := range maps {
		for k, v := range m {
			if _, ok := s[k]; ok {
				s[k+strconv.Itoa(i+1)] = v
			} else {
				s[k] = v
			}
		}
	}
}
