package logger

import (
	"log"
	"os"
)

var LoggerS3 *log.Logger

func init() {
	logPath := "/var/log/app/delivery/s3.log"

	// Open file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// Create logger that writes to the shared file
	LoggerS3 = log.New(file, "[S3] ", log.LstdFlags|log.Lshortfile)
}
