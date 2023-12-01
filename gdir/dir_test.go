package gdir

import (
	"testing"
)

func TestLocateGoModDirectory(t *testing.T) {
	_, err := LocateGoModDirectory()
	if err != nil {
		t.Errorf("Failed to locate GoMod directory: %v", err)
	}
}

func TestFindDirPathOfFileFromGoMod(t *testing.T) {
	filename := "common.go"
	_, err := FindDirPathOfFileFromGoMod(filename)
	if err != nil {
		t.Errorf("Failed to locate file from GoMod directory: %v", err)
	}
}
