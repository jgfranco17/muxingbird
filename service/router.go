package service

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jgfranco17/muxingbird/logging"
)

// Wrapper struct for running the mock HTTP service
type Server struct {
	router *chi.Mux
	port   int
}

// NewRestService creates a new HTTP mock server from a JSON spec.
// It parses the configuration from the provided io.Reader, sets up a router,
// and registers all routes defined in the configuration. The context is used
// to extract a logger for route registration logging.
func NewRestService(ctx context.Context, r io.Reader, port int) (*Server, error) {
	logger := logging.FromContext(ctx)
	serviceConfig, err := LoadFromContent(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %w", err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	for _, e := range serviceConfig.Routes {
		handler := CreateNewMockHandler(e.Status, e.Response)
		router.MethodFunc(e.Method, e.Path, handler)
		logger.Debugf("Registered route: %s %s (returns HTTP %d)", e.Method, e.Path, e.Status)
	}
	logger.Infof("Loaded %d routes", len(serviceConfig.Routes))
	return &Server{
		router: router,
		port:   port,
	}, nil
}

// Run starts the HTTP server and begins listening on the configured port.
func (s *Server) Run(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Infof("Starting service on port %d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
