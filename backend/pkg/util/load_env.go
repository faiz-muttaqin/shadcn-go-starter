package util

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv(envFiles ...string) {
	exePath, _ := os.Executable()
	appDir := filepath.Dir(exePath)
	workDir, _ := os.Getwd()
	envFile := ".env"
	if len(envFiles) > 0 {
		envFile = envFiles[0]
	}
	if err := godotenv.Load(filepath.Join(appDir, envFile)); err != nil {
		appDir = workDir
		if err := godotenv.Load(filepath.Join(appDir, envFile)); err != nil {
			logrus.Errorf("Error loading .env file from both executable directory and current working directory: %v", err)
		}
	}
	os.Setenv("APP_DIR", appDir)
}
