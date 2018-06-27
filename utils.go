package main

import (
	"fmt"
	"os"
	"os/exec"
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

// CloneTckRepo clones raml-tck repo and returns cloned repo path
func CloneTckRepo() string {
	targetDir := fmt.Sprintf("%s/raml-tck", os.TempDir())
	_ = os.RemoveAll(targetDir)
	fmt.Printf("Cloning raml-tc repo to %s\n", targetDir)
	gitRepo := "git@github.com:raml-org/raml-tck.git"
	cmd := exec.Command("git", "clone", gitRepo, targetDir)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to clone repo %s", gitRepo))
	}
	return fmt.Sprintf("%s/tests/raml-1.0", targetDir)
}

// ShouldFail reports whether parsing of RAML file should fail
func ShouldFail(fpath string) bool {
	return strings.Contains(fpath, "invalid")
}
