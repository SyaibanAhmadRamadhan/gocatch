package JOfile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetModuleName() (moduleName string, err error) {
	dir, err := GetGoModPath()
	if err != nil {
		return
	}
	dir += "/go.mod"
	fmt.Println(dir)

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	if strings.HasPrefix(line, "module") {
		moduleName = strings.TrimPrefix(line, "module ")
	} else {
		err = fmt.Errorf("module name not found")
	}

	return
}
func GetGoModPath() (string, error) {
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

func FindFile(filename string) (dir string, err error) {
	rootPath, err := GetGoModPath()
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
