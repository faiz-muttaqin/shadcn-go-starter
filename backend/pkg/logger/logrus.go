package logger

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger initializes the Logrus logger with daily log rotation
func InitLogrus() {
	setLogLevel()
	setLogFormatter()
	fullLogPath := getFullLogPath()

	logrus.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   fullLogPath,
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   true,
	}))

	// logrus.Info("Logger initialized successfully")
}
func setLogLevel() {
	levelMap := map[string]logrus.Level{
		"panic":   logrus.PanicLevel,
		"fatal":   logrus.FatalLevel,
		"error":   logrus.ErrorLevel,
		"warn":    logrus.WarnLevel,
		"warning": logrus.WarnLevel,
		"info":    logrus.InfoLevel,
		"debug":   logrus.DebugLevel,
		"trace":   logrus.TraceLevel,
	}

	envLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel, ok := levelMap[envLevel]; ok {
		logrus.SetLevel(logLevel)
	} else {
		logrus.SetLevel(logrus.TraceLevel) // fallback default
	}
}

func setLogFormatter() {
	formatMap := map[string]logrus.Formatter{
		"text": &logrus.TextFormatter{FullTimestamp: true},
		"json": &logrus.JSONFormatter{},
	}

	format := strings.ToLower(os.Getenv("LOG_FORMAT"))
	formatter, ok := formatMap[format]
	if !ok {
		formatter = &CSVFormatter{
			ForceColors:     true,
			TimestampFormat: " 2006-01-02 15:04:05.000 -07:00",
		}
	}
	logrus.SetFormatter(formatter)
	logrus.SetReportCaller(true)
}

func getFullLogPath() string {
	logPath := os.Getenv("LOG_PATH")
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "app.log"
	}
	if logPath == "" {
		logPath = "./log/app/"
	}
	if strings.HasPrefix(logPath, "./") {
		appDir := os.Getenv("APP_DIR")
		logPath = filepath.Join(appDir, logPath[2:])
	}
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		_ = os.MkdirAll(logPath, os.ModePerm)
	}
	return filepath.Join(logPath, logFile)
}
