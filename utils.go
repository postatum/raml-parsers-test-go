package main

import (
	"os"
	"path/filepath"
	"strings"
)

// ListRamls lists RAML file in :folderPath: folder
func ListRamls(folderPath string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(folderPath, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".raml") {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
