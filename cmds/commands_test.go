package cmds

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/jgfranco17/muxingbird/internal"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	runCalled   bool
	expectedErr error
}

func (m *mockService) Run(ctx context.Context) error {
	m.runCalled = true
	return m.expectedErr
}

func (m *mockService) WasCalled() bool {
	return m.runCalled
}

func newMockFactory(service *mockService) ServiceFactory {
	return func(ctx context.Context, r io.Reader, port int) (HttpService, error) {
		return service, nil
	}
}

func applyContextLogger(cmd *cobra.Command) {
	logger := logging.NewLogger()
	ctx := logging.ApplyToContext(context.Background(), logger)
	cmd.SetContext(ctx)
}

func TestRunCommandSuccess(t *testing.T) {
	mock := &mockService{}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	applyContextLogger(cmd)
	result := internal.ExecuteTestCommand(t, cmd, "./resources/mock.json")
	assert.NoError(t, result.Error, "Unexpected error while executing run command")
}

func TestRunCommandFail_InvalidPath(t *testing.T) {
	mock := &mockService{}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	applyContextLogger(cmd)
	result := internal.ExecuteTestCommand(t, cmd, "nonexistent")
	assert.ErrorContains(t, result.Error, "no such file")
}

func TestRunCommandFail_ServiceFactoryError(t *testing.T) {
	mock := &mockService{
		expectedErr: errors.New("some mock error"),
	}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	applyContextLogger(cmd)
	result := internal.ExecuteTestCommand(t, cmd, "./resources/mock.json")
	assert.ErrorContains(t, result.Error, "some mock error")
}
