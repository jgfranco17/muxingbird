package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

const (
	ConfigFileName string = "muxingbird.yaml"
)

// ServerConfig represents a single HTTP server's configuration.
type ServerConfig struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type MockResponseJson map[string]interface{}

// Route represents an individual route with method, path, and static response.
type Route struct {
	Path     string           `json:"path"`
	Method   string           `json:"method"`
	Status   int              `json:"status"`
	Response MockResponseJson `json:"response"`
}

// LoadFromContent reads and parses a JSON content into a MockServer.
func LoadFromContent(r io.Reader) (*ServerConfig, error) {
	var server ServerConfig
	if err := yaml.NewDecoder(r).Decode(&server); err != nil {
		return nil, fmt.Errorf("invalid config format: %w", err)
	}
	return &server, nil
}

// InitConfig writes a sample mock server configuration to the provided writer.
// The generated configuration includes a basic HTTP GET route with a sample
// JSON response. This function is intended to help users bootstrap a valid
// configuration file.
func InitConfig(ctx context.Context, w io.Writer) error {
	example := &ServerConfig{
		Name: "new-mock-server",
		Routes: []Route{
			{
				Method:   "GET",
				Path:     "/hello",
				Status:   200,
				Response: MockResponseJson{"message": "Hello, world!"},
			},
		},
	}
	prettyPrintJson, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode sample config: %w", err)
	}
	if _, err = fmt.Fprintln(w, string(prettyPrintJson)); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return nil
}
