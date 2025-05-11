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

const (
	localhostAddr string = "127.0.0.1"
)

// Wrapper struct for running the mock HTTP service
type Server struct {
	router *chi.Mux
	port   int
	config *ServerConfig
}

// NewRestService creates a new HTTP mock server from a JSON spec.
// It parses the configuration from the provided io.Reader, sets up a router,
// and registers all routes defined in the configuration. The context is used
// to extract a logger for route registration logging.
func NewRestService(ctx context.Context, r io.Reader, port int) (*Server, error) {
	logger := logging.FromContext(ctx)
	cfg, err := LoadFromContent(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	for _, e := range cfg.Routes {
		handler := CreateNewMockHandler(e.Status, e.Response)
		router.MethodFunc(e.Method, e.Path, handler)
		logger.Tracef("Registered route: %s %s (returns HTTP %d)", e.Method, e.Path, e.Status)
	}
	logger.Tracef("Loaded %d routes from specification", len(cfg.Routes))
	return &Server{
		router: router,
		port:   port,
		config: cfg,
	}, nil
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", localhostAddr, s.port)
}

// Run starts the HTTP server and begins listening on the configured port.
func (s *Server) Run(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Infof("Starting service at http://%s", s.Address())
	return http.ListenAndServe(s.Address(), s.router)
}
