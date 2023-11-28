package JOFD

import (
	"testing"
)

func TestGetModuleName(t *testing.T) {
	moduleName, err := GetModuleName()
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	t.Log(moduleName)
}
