package gdir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindDirPathOfFileFromGoMod finds the relative path of the directory containing the filename in a Go module.
// It first locates the Go module directory, then it traverses the directory tree within this module.
// If it finds the file, it returns the relative path of its directory. If it fails, it returns an error.
func FindDirPathOfFileFromGoMod(filename string) (dir string, err error) {
	rootPath, err := LocateGoModDirectory()
	if err != nil {
		return
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failure accessing a path %q: %v\n", path, err)
		}

		// If the info points to a file and so it has the same name as the filename, get the relative path of the directory
		if !info.IsDir() && strings.HasSuffix(info.Name(), filename) {
			dir, err = filepath.Rel(rootPath, filepath.Dir(path))
			if err != nil {
				return fmt.Errorf("failure getting relative path for %s: %v\n", path, err)
			}
		}

		return nil
	})

	return
}

// LocateGoModDirectory finds the directory containing the go.mod file, starting from the current working directory and moving up.
// It returns the path of the directory or an error if it's not found.
func LocateGoModDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}
