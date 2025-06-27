package logger

import (
	"io"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var LoggerS3 *log.Logger

func init() {
	logPath := "/var/log/app/delivery/s3.log"

	// Set up lumberjack for rotation
	rotator := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
		Compress:   true,
	}

	// MultiWriter for both stdout and file
	multi := io.MultiWriter(rotator)
	LoggerS3 = log.New(multi, "[S3] ", log.LstdFlags|log.Lshortfile)
}
