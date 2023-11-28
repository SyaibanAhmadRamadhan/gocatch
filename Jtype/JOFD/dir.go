package JOFD

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindDirPathFileFromGoModule(filename string) (dir string, err error) {
	rootPath, err := FindGoModPath()
	if err != nil {
		return
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failure accessing a path %q: %v\n", path, err)
		}

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

func FindGoModPath() (string, error) {
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
