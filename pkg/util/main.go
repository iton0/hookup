package util

import (
	"os"
	"path/filepath"
	"strings"
)

// Defaults to cwd if path not given
func CreateFolder(path []string, name string) (string, error) {
	dirPath := toPathString(path, name)
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return "", err
	}
	return dirPath, nil
}

// Defaults to cwd if path not given
func CreateFile(path []string, name string) (string, error) {
	filePath := toPathString(path, name)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close() // No memory leaks here no sir no sir
	return filePath, nil
}

// Defaults to cwd if path not given
func DoesDirectoryExist(path []string) bool {
	info, err := os.Stat(toPathString(path))
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// SplitFilePath takes a string representation of a file path and returns a slice representation of it.
func SplitFilePath(path string) (slice []string) {
	slice = strings.Split(path, string(filepath.Separator))
	return // return slice

}

// toPathString combines a slice of path components and an optional name into a single path string.
// If no name is provided, it joins the components in the path slice.
// If a name is provided, it appends the name to the path slice before joining.
// The resulting path uses the appropriate file separators for the operating system.
func toPathString(path []string, name ...string) string {
	if len(name) == 0 {
		return filepath.Join(path...)
	}
	return filepath.Join(append(path, name...)...)
}
