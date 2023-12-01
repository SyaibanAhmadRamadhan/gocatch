package gcommon

import (
	"github.com/oklog/ulid/v2"
)

// NewUlid generates a new ULID as a string
func NewUlid() string {
	return ulid.Make().String()
}

// ParseUlid parses the ULID from a string, with an option for 'strict' mode
// The 'strict' mode will only allow valid ULID strings
func ParseUlid(s string, strict bool) (id ulid.ULID, err error) {
	if strict {
		return ulid.ParseStrict(s)
	}
	return ulid.Parse(s)
}
