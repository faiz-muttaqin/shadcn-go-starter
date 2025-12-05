package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataDir(appName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var base string

	switch runtime.GOOS {
	case "windows":
		// Example: C:\Users\User\AppData\Local\MyApp\data
		appData := os.Getenv("LOCALAPPDATA")
		if appData == "" {
			appData = filepath.Join(home, "AppData", "Local")
		}
		base = filepath.Join(appData, appName)

	case "darwin":
		// macOS: /Users/User/Library/Application Support/MyApp/data
		base = filepath.Join(home, "Library", "Application Support", appName)

	case "android":
		// Termux / Android: /data/data/com.termux/files/home/.local/share/MyApp/data
		base = filepath.Join(home, ".local", "share", appName)

	default:
		// Linux / Ubuntu: /home/user/.local/share/MyApp/data
		base = filepath.Join(home, ".local", "share", appName)
	}

	dataPath := filepath.Join(base, "data")

	// ensure directory exists
	if err := os.MkdirAll(dataPath, 0755); err != nil {
		return "", err
	}

	return dataPath, nil
}
