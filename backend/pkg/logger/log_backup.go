package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func LogBackup(logDir string) {
	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}

	// Get the current date in the format "YYYY_MM_DD"
	currentDate := time.Now().AddDate(0, 0, -1).Format("2006_01_02")

	// Path to the source log file
	sourceLogFilePath := filepath.Join(logDir, "apps.log")

	// Path to the target log file with today's date
	targetLogFilePath := filepath.Join(logDir, fmt.Sprintf("apps_%s.log", currentDate))

	// Check if the target log file for today's date exists
	if _, err := os.Stat(targetLogFilePath); os.IsNotExist(err) {
		// Target log file doesn't exist, create it

		// Read the content of the source log file
		content, err := os.ReadFile(sourceLogFilePath)
		if err != nil {
			logrus.Error(err)
			logrus.Fatal(err)
		}

		// Write the content to the target log file
		err = os.WriteFile(targetLogFilePath, content, os.ModePerm)
		if err != nil {
			logrus.Error(err)
			logrus.Fatal(err)
		}

		// Empty the source log file
		err = os.WriteFile(sourceLogFilePath, nil, os.ModePerm)
		if err != nil {
			logrus.Error(err)
			logrus.Fatal(err)
		}

		message := fmt.Sprintf("Log files Backup %s Done", currentDate)
		fmt.Println(message)
	}
}
