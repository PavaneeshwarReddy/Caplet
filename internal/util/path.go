package util

import (
	"os"
	"path/filepath"
)

func GetWorkingDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "caplet"), nil
}

func GetConfigPath(workingDir string) string {
	return filepath.Join(workingDir, "config.json")
}
