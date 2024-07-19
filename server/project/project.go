package project

import (
	"os"
	"path/filepath"
)

var root string

// Root returns healthcare-server project root absolute path
func Root() string {
	if root != "" {
		return root
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		path := filepath.Clean(filepath.Join(currentDir, "go.mod"))
		_, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	root = currentDir

	return root
}
