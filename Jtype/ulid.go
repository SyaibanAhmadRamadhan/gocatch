package Jtype

import (
	"github.com/oklog/ulid/v2"
)

func NewUlid() string {
	return ulid.Make().String()
}

func ParseUlid(s string, strict bool) (id ulid.ULID, err error) {
	if strict {
		return ulid.ParseStrict(s)
	}
	return ulid.Parse(s)
}
