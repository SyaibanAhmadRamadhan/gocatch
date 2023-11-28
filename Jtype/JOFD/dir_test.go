package JOFD

import (
	"testing"
)

func TestGetGoModPath(t *testing.T) {
	dir, err := FindGoModPath()
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	t.Log(dir)
}

func TestFindDirPathFileFromGoModule(t *testing.T) {
	dir, err := FindDirPathFileFromGoModule("postgres_generator.go")
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	t.Log(dir)
}
