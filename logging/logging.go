package logging

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// NewLogger configures and registers a new logger instance.
func NewLogger() *logrus.Logger {
	logger := logrus.StandardLogger()
	logger.SetReportCaller(true)
	logger.SetFormatter(&CustomFormatter{})
	return logger
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Create a custom format for the log message
	level := strings.ToUpper(entry.Level.String())
	timestamp := entry.Time.Format(time.DateTime)
	colorFunc := color.New(setOutputColorPerLevel(level)).SprintFunc()
	logMessage := fmt.Sprintf("[%s][%s] %s\n", timestamp, colorFunc(level), entry.Message)
	return []byte(logMessage), nil
}

func setOutputColorPerLevel(level string) color.Attribute {
	var selectedColor color.Attribute
	switch level {
	case "DEBUG":
		selectedColor = color.FgCyan
	case "INFO":
		selectedColor = color.FgGreen
	case "WARN", "WARNING":
		selectedColor = color.FgYellow
	case "ERROR", "PANIC", "FATAL":
		selectedColor = color.FgRed
	default:
		selectedColor = color.FgWhite
	}
	return selectedColor
}
