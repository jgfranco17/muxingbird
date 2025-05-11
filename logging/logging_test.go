package logging

import (
	"context"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomFormatterFormat(t *testing.T) {
	// Create a CustomFormatter instance
	formatter := &CustomFormatter{}

	// Create a logrus Entry
	entry := &logrus.Entry{
		Logger:  logrus.New(),
		Data:    logrus.Fields{},
		Time:    time.Now(),
		Level:   logrus.InfoLevel,
		Message: "This is a test log message",
	}

	// Format the entry
	output, err := formatter.Format(entry)
	assert.NoError(t, err)
	outputStr := string(output)
	assert.Contains(t, outputStr, "INFO")
	assert.Contains(t, outputStr, entry.Message)

	expectedTimestamp := entry.Time.Format(time.TimeOnly)
	assert.Contains(t, outputStr, expectedTimestamp)
	colorFunc := color.New(color.FgGreen).SprintFunc()
	assert.Contains(t, outputStr, colorFunc("INFO"))
}

func TestApplyToContextAndFromContext_Success(t *testing.T) {
	ctx := context.Background()
	logger := logrus.New()

	ctxWithLogger := ApplyToContext(ctx, logger)

	retrieved := FromContext(ctxWithLogger)
	require.NotNil(t, retrieved, "Logger should not be nil")
	assert.Equal(t, logger, retrieved, "Retrieved logger should match the original")
}

func TestFromContext_PanicsWhenMissing(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Expected panic when logger is missing from context")
	}()

	ctx := context.TODO()
	_ = FromContext(ctx)
}
