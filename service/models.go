package service

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
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
		return nil, fmt.Errorf("Invalid config format: %w", err)
	}
	return &server, nil
}
