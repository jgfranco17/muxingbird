package cmds

import (
	"context"
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/jgfranco17/muxingbird/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testValidSpecJsonFile string = "./resources/mock.json"
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

func TestRunCommandSuccess(t *testing.T) {
	mock := &mockService{}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	result := internal.ExecuteTestCommand(t, cmd, "-f", testValidSpecJsonFile)
	assert.NoError(t, result.Error, "Unexpected error while executing run command")
}

func TestRunCommandFail_InvalidPath(t *testing.T) {
	mock := &mockService{}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	result := internal.ExecuteTestCommand(t, cmd, "-f", "nonexistent")
	assert.ErrorContains(t, result.Error, "no such file")
}

func TestRunCommandFail_ServiceFactoryError(t *testing.T) {
	mock := &mockService{
		expectedErr: errors.New("some mock error"),
	}
	factory := newMockFactory(mock)
	cmd := CommandRun(factory)
	result := internal.ExecuteTestCommand(t, cmd, "-f", testValidSpecJsonFile)
	assert.ErrorContains(t, result.Error, "some mock error")
}

func TestInitCommandSuccess_StdoutDefault(t *testing.T) {
	result := internal.ExecuteTestCommand(t, CommandInit())
	require.NoError(t, result.Error, "Unexpected error while running init command to stdout")
}

func TestInitCommandSuccess_CreateFile(t *testing.T) {
	sampleInitOutputFile := "init-test.json"
	mockOutputPath := filepath.Join(t.TempDir(), sampleInitOutputFile)
	t.Run("new file", func(t *testing.T) {
		resultToStdout := internal.ExecuteTestCommand(t, CommandInit(), "-o", mockOutputPath)
		require.NoError(t, resultToStdout.Error, "Unexpected error while running init command with file")
	})
	t.Run("existing file", func(t *testing.T) {
		resultToStdout := internal.ExecuteTestCommand(t, CommandInit(), "-o", mockOutputPath, "--force")
		require.NoError(t, resultToStdout.Error, "Unexpected error while running init command with force")
	})
}

func TestInitCommandFail_NoForceButFileExists(t *testing.T) {
	result := internal.ExecuteTestCommand(t, CommandInit(), "-o", testValidSpecJsonFile)
	require.Error(t, result.Error, "Wanted error, but got none")
}
