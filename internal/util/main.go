package util

import (
	"os"
)

// Defaults to cwd if path not given
func CreateFolder(dirPath string) error {
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// Defaults to cwd if path not given
func CreateFile(dirPath string, name string) error {
	filePath := dirPath + name
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close() // No memory leaks here no sir no sir
	return nil
}

// Defaults to cwd if path not given
func DoesDirectoryExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
