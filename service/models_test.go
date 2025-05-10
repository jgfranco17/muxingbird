package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFromContent_ParserSuccess(t *testing.T) {
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
        "error": "Invalid input error"
      }
    }
  ]
}`)
	server, err := LoadFromContent(mockContent)
	require.NoError(t, err, "Unexpected error loading from valid config")

	expectedConfig := &ServerConfig{
		Name: "example-mock-server",
		Routes: []Route{
			{
				Path:   "/hello",
				Method: "GET",
				Status: 200,
				Response: MockResponseJson{
					"message": "Hello, world!",
				},
			},
			{
				Path:   "/status",
				Method: "POST",
				Status: 201,
				Response: MockResponseJson{
					"message": "Status created successfully",
				},
			},
			{
				Path:   "/error",
				Method: "GET",
				Status: 400,
				Response: MockResponseJson{
					"error": "Invalid input error",
				},
			},
		},
	}
	assert.Equal(t, expectedConfig, server)
}

func TestLoadFromContentFail_InvalidContentFormat(t *testing.T) {
	mockContent := strings.NewReader("foo bar")
	_, err := LoadFromContent(mockContent)
	assert.ErrorContains(t, err, "invalid config format")
}
