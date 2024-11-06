package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigDir() (string, error) {

	var basePath string

	switch runtime.GOOS {
	case "windows":

		basePath = os.Getenv("APPDATA")

		if basePath == "" {
			return "", fmt.Errorf("environment variable is not set: APPDATA")
		}

	case "linux", "darwin":

		homeDir, err := os.UserHomeDir()

		if err != nil {
			return "", fmt.Errorf("could not get user home directory: %w", err)
		}

		basePath = filepath.Join(homeDir, ".config")

	default:
		return "", fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}

	configDir := filepath.Join(basePath, "jencli")

	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create configuration directory: %w", err)
	}

	return configDir, nil

}
