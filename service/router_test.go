package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jgfranco17/muxingbird/internal"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock configuration JSON to be used in tests
const mockConfig = `{
	"routes": [
		{
			"method": "GET",
			"path": "/hello",
			"status": 200,
			"response": {"message": "Hello, world!"}
		},
		{
			"method": "POST",
			"path": "/goodbye",
			"status": 201,
			"response": {"message": "Goodbye, world!"}
		}
	]
}`

// Test NewRestService with valid configuration
func TestNewRestService_Success(t *testing.T) {
	mockReader := bytes.NewReader([]byte(mockConfig))
	ctx := logging.ApplyToContext(context.Background(), internal.NewTestLogger(t))
	service, err := NewRestService(ctx, mockReader, 8080)

	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, 8080, service.port)
	assert.Len(t, service.config.Routes, 2)
}

// Test NewRestService with invalid configuration (invalid JSON)
func TestNewRestService_InvalidConfig(t *testing.T) {
	mockReader := bytes.NewReader([]byte(`{"invalidJson`))
	ctx := logging.ApplyToContext(context.Background(), internal.NewTestLogger(t))
	service, err := NewRestService(ctx, mockReader, 8080)
	require.Error(t, err)
	assert.Nil(t, service)
	assert.Contains(t, err.Error(), "failed to read config")
}

// Test Server Address function
func TestServer_Address(t *testing.T) {
	service := &Server{
		port: 8080,
	}
	address := service.Address()
	assert.Equal(t, "127.0.0.1:8080", address)
}

// Test Server Run method with a mock HTTP request
func TestServer_Run(t *testing.T) {
	mockReader := bytes.NewReader([]byte(mockConfig))
	ctx := logging.ApplyToContext(context.Background(), internal.NewTestLogger(t))
	service, err := NewRestService(ctx, mockReader, 8080)
	assert.NoError(t, err)

	// Create a test server to simulate running the service
	ts := httptest.NewServer(service.router)
	defer ts.Close()

	// Make a GET request to /hello route
	resp, err := http.Get(fmt.Sprintf("%s/hello", ts.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	var respBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, world!", respBody["message"])

	// Test POST route
	resp, err = http.Post(fmt.Sprintf("%s/goodbye", ts.URL), "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

// Test middleware logging setup (check for logger usage in the service setup)
func TestMiddleware_Logger(t *testing.T) {
	mockReader := bytes.NewReader([]byte(mockConfig))
	ctx := logging.ApplyToContext(context.Background(), internal.NewTestLogger(t))
	service, err := NewRestService(ctx, mockReader, 8080)
	require.NoError(t, err)
	assert.NotNil(t, service)

	// Test logger middleware
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/hello", nil)
	assert.NoError(t, err)

	// Create middleware chain and invoke
	service.router.ServeHTTP(w, req)

	// Check if the response status is what we expect
	assert.Equal(t, http.StatusOK, w.Code)
}
