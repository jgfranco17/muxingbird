package cmds

import (
	"bytes"
	"context"
	"testing"

	"github.com/jgfranco17/muxingbird/internal"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultServiceFactory(t *testing.T) {
	// Mock configuration JSON to be used in tests
	const mockConfig = `{
	"routes": [
		{
			"method": "GET",
			"path": "/hello",
			"status": 200,
			"response": {"message": "Hello, world!"}
		},
	]
}`
	mockReader := bytes.NewReader([]byte(mockConfig))
	ctx := logging.ApplyToContext(context.Background(), internal.NewTestLogger(t))
	srv, err := DefaultServiceFactory(ctx, mockReader, 9000)
	require.NoError(t, err)
	assert.NotNil(t, srv)
}
