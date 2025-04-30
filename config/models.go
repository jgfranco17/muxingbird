package config

// MockServer represents a single HTTP server's configuration.
type MockServer struct {
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
