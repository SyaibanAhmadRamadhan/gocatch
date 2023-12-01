package gcommon

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestNewUlid(t *testing.T) {
	result := NewUlid()
	if len(result) != 26 {
		t.Errorf("Expected length 26, but got %v", len(result))
	}
	if _, err := ulid.Parse(result); err != nil {
		t.Errorf("Failed parsing ULID: %v", err)
	}
}

func TestParseUlid(t *testing.T) {
	ulidStr := NewUlid()
	_, err := ParseUlid(ulidStr, true)
	if err != nil {
		t.Errorf("Failed parsing ULID: %v", err)
	}

	_, err = ParseUlid(ulidStr, false)
	if err != nil {
		t.Errorf("Failed parsing ULID: %v", err)
	}
}
