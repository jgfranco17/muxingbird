package internal

import (
	"bytes"
	"context"
	"testing"

	"github.com/jgfranco17/muxingbird/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CliCommandFunction func() *cobra.Command

type CommandRunner func(cmd *cobra.Command, args []string)

type CliRunResult struct {
	Output string
	Error  error
}

// Helper function to simulate CLI execution
func ExecuteTestCommand(t *testing.T, cmd *cobra.Command, args ...string) CliRunResult {
	t.Helper()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	logger := NewTestLogger(t)
	ctx := logging.ApplyToContext(context.Background(), logger)
	cmd.SetContext(ctx)

	_, err := cmd.ExecuteC()
	return CliRunResult{
		Output: buf.String(),
		Error:  err,
	}
}

func NewTestLogger(t *testing.T) *logrus.Logger {
	t.Helper()
	logger := logrus.New()
	logger.SetOutput(new(bytes.Buffer))
	logger.Level = logrus.DebugLevel
	logger.SetReportCaller(true)
	logger.SetFormatter(new(logrus.TextFormatter))
	return logger
}
