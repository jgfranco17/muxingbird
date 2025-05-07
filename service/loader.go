package service

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// LoadFromContent reads and parses a JSON content into a MockServer.
func LoadFromContent(r io.Reader) (*MockServer, error) {
	var server MockServer
	if err := yaml.NewDecoder(r).Decode(&server); err != nil {
		return nil, fmt.Errorf("Invalid config format: %w", err)
	}
	return &server, nil
}
