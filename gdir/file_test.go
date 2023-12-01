package gdir

import (
	"testing"
)

func TestGetModuleName(t *testing.T) {
	moduleName, err := GetModuleName()
	if err != nil {
		t.Errorf("Failed to get module name: %v", err)
	}
	if moduleName == "" {
		t.Errorf("Expected a module name, got empty string")
	}
}
