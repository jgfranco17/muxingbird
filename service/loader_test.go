package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromContentSuccess(t *testing.T) {
	mockContent := strings.NewReader(`{
  "name": "example-mock-server",
  "routes": [
    {
      "path": "/hello",
      "method": "GET",
      "status": 200,
      "response": {
        "message": "Hello, world!"
      }
    },
    {
      "path": "/status",
      "method": "POST",
      "status": 201,
      "response": {
        "message": "Status created successfully"
      }
    },
    {
      "path": "/error",
      "method": "GET",
      "status": 400,
      "response": {
        "error": "Internal server error"
      }
    }
  ]
}`)
	server, err := LoadFromContent(mockContent)
	assert.NoError(t, err, "Unexpected error loading from valid config")
	assert.Equal(t, "example-mock-server", server.Name)
	assert.Lenf(t, server.Routes, 3, "Expected 3 routes but got %d", len(server.Routes))
}

func TestLoadFromContentFail(t *testing.T) {
	mockContent := strings.NewReader("foo bar")
	_, err := LoadFromContent(mockContent)
	assert.ErrorContains(t, err, "Invalid config format")
}
