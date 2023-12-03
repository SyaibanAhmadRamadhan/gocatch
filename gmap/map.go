package gmap

import (
	"strconv"
	"strings"
)

// StrAny is a type alias for a map of string keys to any values.
type StrAny map[string]any

// ConcatKeys takes prefix and separator as parameters and concatenates all
// keys of the map. If the map is nil, an empty string is returned. Prefix defaults to "".
// Separator defaults to ", ".
// [0] IS PREFIX.
// [1] IS SEPARATOR.
func (s StrAny) ConcatKeys(ps ...string) string {
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
	for k := range s {
		saArr = append(saArr, prefix+k)
	}

	return strings.Join(saArr, separator)
}

// Merge takes one or more map[string]any parameters and adds their entries to
// the original map. If the key already exists, an integer index is appended.
func (s StrAny) Merge(maps ...map[string]any) {
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
