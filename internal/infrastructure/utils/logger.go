package utils

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	base *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		base: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Error(message string, err error, fatal bool) {
	l.base.Printf("[ERROR] [%s] %s : %v",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
		err,
	)

	l.base.Printf("[ERROR] [%s] %s : %v",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
		err,
	)

	if fatal {
		l.base.Fatalf("[FATAL] [%s] %s : %v",
			time.Now().Format("2006-01-02 15:04:05"),
			message,
			err,
		)
	}
}

func (l *Logger) Info(message string) {
	l.base.Printf("[INFO] [%s] %s",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
	)
}
