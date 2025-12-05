package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/sirupsen/logrus"
)

// CustomResponseWriter captures response body
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures response body
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Save response body
	return w.ResponseWriter.Write(b)
}

var (
	LogRoutes     bool
	LogReqHeaders bool
	LogReqBody    bool
	LogResHeaders bool
	LogResBody    bool
)

// Must be deferred after opening log file
func InitGinLogger() (*os.File, error) {
	LogRoutes = util.Getenv("LOG_ROUTES", false)
	LogReqHeaders = util.Getenv("LOG_REQ_HEADER", false)
	LogReqBody = util.Getenv("LOG_REQ_BODY", false)
	LogResHeaders = util.Getenv("LOG_RES_HEADER", false)
	LogResBody = util.Getenv("LOG_RES_BODY", false)

	logPath := util.Getenv("LOG_ROUTES_PATH", "./log/routes/")
	logFile := util.Getenv("LOG_ROUTES_FILE", "routes.log")

	// Normalize path if starts with "./"
	if strings.HasPrefix(logPath, "./") {
		logPath = filepath.Join(os.Getenv("APP_DIR"), logPath[2:])
	}

	logFilePath := filepath.Join(logPath, logFile)

	// 1️⃣ Check and create folder if not exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log folder: %v", err)
		}
	}

	// 2️⃣ Check and create file if not exist
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		f, err := os.Create(logFilePath)
		if err != nil {
			log.Fatalf("Failed to create log file: %v", err)
		}
		defer f.Close()
	}
	osLogFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	go BackupGinLogService(logPath, 30*time.Minute)

	return osLogFile, nil
}

// GinLoggerMiddleware logs request and response
func GinLoggerMiddleware(logFile *os.File) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var reqBodyBytes []byte
		respBodyBuffer := new(bytes.Buffer)
		if LogRoutes {
			// Capture request body
			if LogReqBody && c.Request.Body != nil {
				reqBodyBytes, _ = io.ReadAll(c.Request.Body)
				c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes)) // Restore body
			}

			// Create custom response writer
			customWriter := &CustomResponseWriter{body: respBodyBuffer, ResponseWriter: c.Writer}
			c.Writer = customWriter // Replace Gin's response writer
		}
		// Process request
		c.Next()

		// Log details
		if LogRoutes {
			additional_info := ""
			if LogReqHeaders {
				additional_info = fmt.Sprintf("%s[REQ_HEADER] \n%v\n", additional_info, c.Request.Header)
			}
			if LogReqBody {
				additional_info = fmt.Sprintf("%s[REQ_BODY] \n%s\n", additional_info, string(reqBodyBytes))
			}
			if LogResHeaders {
				additional_info = fmt.Sprintf("%s[RES_HEADER] \n%v\n", additional_info, c.Writer.Header())
			}
			if LogResBody {
				additional_info = fmt.Sprintf("%s[RES_BODY] \n%v\n", additional_info, respBodyBuffer.String())
			}
			userAgent := c.Request.UserAgent()
			ua := user_agent.New(userAgent)
			browser, version := ua.Browser()
			os := ua.OS()
			acessUsername := "?"
			fmt.Fprintf(logFile, "[GIN] %v | %-7s | %3d | %13v | %15s | %10s | %-7s %-9s | %s | %s \n%s",
				start.Format("2006/01/02 - 15:04:05"),
				c.Request.Method,
				c.Writer.Status(),
				time.Since(start),
				c.ClientIP(),
				os,
				browser,
				version,
				acessUsername,
				c.Request.URL.Path,
				additional_info,
			)
		}
	}
}
func BackupGinLogService(appLogDir string, interval time.Duration) {
	//CHECK ONE TIME
	checkLogWrite(appLogDir)

	// Create a ticker for the interval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Call the function you want to execute at each interval
			checkLogWrite(appLogDir)
		}
	}
}
func checkLogWrite(logDir string) {
	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Get the current date in the format "YYYY_MM_DD"
	currentDate := time.Now().AddDate(0, 0, -1).Format("2006_01_02")

	// Path to the source log file
	sourceLogFilePath := filepath.Join(logDir, "routes.log")

	// Path to the target log file with today's date
	targetLogFilePath := filepath.Join(logDir, fmt.Sprintf("routes_%s.log", currentDate))

	// Check if the target log file for today's date exists
	if _, err := os.Stat(targetLogFilePath); os.IsNotExist(err) {
		// Target log file doesn't exist, create it

		// Read the content of the source log file
		content, err := os.ReadFile(sourceLogFilePath)
		if err != nil {
			logrus.Error(err)
			log.Fatal(err)
		}

		// Write the content to the target log file
		err = os.WriteFile(targetLogFilePath, content, os.ModePerm)
		if err != nil {
			logrus.Error(err)
			log.Fatal(err)
		}

		// Empty the source log file
		err = os.WriteFile(sourceLogFilePath, nil, os.ModePerm)
		if err != nil {
			logrus.Error(err)
			log.Fatal(err)
		}

		message := fmt.Sprintf("Log files Backup %s Done", currentDate)
		fmt.Println(message)
	}
}
