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
	timestamp := entry.Time.Format(time.DateTime)
	colorFunc := color.New(setOutputColorPerLevel(entry.Level)).SprintFunc()
	logMessage := fmt.Sprintf("[%s][%s] %s\n", timestamp, colorFunc(strings.ToUpper(entry.Level.String())), entry.Message)
	return []byte(logMessage), nil
}

func setOutputColorPerLevel(level logrus.Level) color.Attribute {
	var selectedColor color.Attribute
	switch level {
	case logrus.TraceLevel:
		selectedColor = color.FgBlue
	case logrus.DebugLevel:
		selectedColor = color.FgCyan
	case logrus.InfoLevel:
		selectedColor = color.FgGreen
	case logrus.WarnLevel:
		selectedColor = color.FgYellow
	case logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel:
		selectedColor = color.FgRed
	default:
		selectedColor = color.FgWhite
	}
	return selectedColor
}
