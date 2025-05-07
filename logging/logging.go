package logging

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type ctxKeyLogger struct{}

// NewLogger configures and registers a new logger instance.
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&CustomFormatter{})
	return logger
}

// ApplyToContext attaches a logger to the given context
func ApplyToContext(ctx context.Context, logger *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger{}, logger)
}

// FromContext loads a logger from context. Any code path MUST be
// sure that a logger is attached, or the function will panic.
func FromContext(ctx context.Context) *logrus.Logger {
	logger, ok := ctx.Value(ctxKeyLogger{}).(*logrus.Logger)
	if !ok {
		panic("No logger attached in context")
	}
	return logger
}

type CustomFormatter struct{}

// Format the log entry into clean, colored log messages.
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(time.TimeOnly)
	colorFunc := color.New(setOutputColorPerLevel(entry.Level), color.Bold).SprintFunc()
	coloredLevel := colorFunc(strings.ToUpper(entry.Level.String()))
	logMessage := fmt.Sprintf("[%s][%s] %s\n", timestamp, coloredLevel, entry.Message)
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
