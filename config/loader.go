package config

import (
	"encoding/json"
	"fmt"
	"io"
)

// LoadFromReader reads and parses a JSON content into a MockServer.
func LoadFromReader(r io.Reader) (*MockServer, error) {
	var server MockServer
	if err := json.NewDecoder(r).Decode(&server); err != nil {
		return nil, fmt.Errorf("Invalid config format: %w", err)
	}
	return &server, nil
}
