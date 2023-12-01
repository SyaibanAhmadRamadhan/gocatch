package gdir

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetModuleName reads the go.mod file in the project root,
// extracts the module name from it, and returns the module name.
// If the module name is not found, or an error occurs, it returns an error.
func GetModuleName() (moduleName string, err error) {
	dir, err := LocateGoModDirectory()
	if err != nil {
		return
	}
	dir += "/go.mod"

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
