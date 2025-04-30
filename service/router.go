package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jgfranco17/muxingbird/config"
	"github.com/jgfranco17/muxingbird/logging"
)

type Server struct {
	router *chi.Mux
	port   int
}

func NewRestService(r io.Reader, port int) (*Server, error) {
	logger := logging.NewLogger()
	serviceConfig, err := config.LoadFromReader(r)
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
	}
	logger.Debugf("Loaded %d routes", len(serviceConfig.Routes))
	return &Server{
		router: router,
		port:   port,
	}, nil
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
