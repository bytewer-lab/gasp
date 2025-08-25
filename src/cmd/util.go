package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindGitRoot() (string, error) {
	startDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %s\n", err)
		os.Exit(1)
	}
	dir := startDir
	for {
		_, err := os.Stat(filepath.Join(dir, ".git"))
		if err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir { // Reached the root directory
			return "", fmt.Errorf("no git repository found in current directory or any parent directories")
		}
		dir = parent
	}
}

func CleanTargetDirectory(gitHooksDir string) error {
	filesInTargetDirectory, err := os.ReadDir(gitHooksDir)
	if err != nil {
		return fmt.Errorf("failed to read hooks directory: %s", err)
	}
	for _, file := range filesInTargetDirectory {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if !strings.HasSuffix(filename, ".sample") {
			err = os.Remove(filepath.Join(gitHooksDir, filename))
			if err != nil {
				return fmt.Errorf("failed to remove hook file: %s", filename)
			}
		}
	}
	return nil
}
