package gcommon

import (
	"testing"
)

func TestRandomAlphabeticString(t *testing.T) {
	result := RandomAlphabeticString(5)
	if len(result) != 5 {
		t.Errorf("Expected length 5, but got %v", len(result))
	}
}

func TestRandomStringFromCharset(t *testing.T) {
	customCharset := []rune("0123456789")
	result := RandomStringFromCharset(5, customCharset)
	if len(result) != 5 {
		t.Errorf("Expected length 5, but got %v", len(result))
	}
}

func TestRandomNumericString(t *testing.T) {
	result := RandomNumericString(6)
	if len(result) != 6 {
		t.Errorf("Expected length 6, but got %v", len(result))
	}
}
