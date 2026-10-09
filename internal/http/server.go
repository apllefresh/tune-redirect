package http

import (
	"net/http"
	"os"

	"github.com/apllefresh/tune-redirect/internal/handler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func NewServer() *Server {
	httpAddress := os.Getenv("HTTP_ADDR")
	if httpAddress == "" {
		panic("HTTP_ADDR is required")
	}

	router := chi.NewRouter()

	s := &http.Server{
		Addr:    httpAddress,
		Handler: router,
	}

	return &Server{
		httpServer: s,
		router:     router,
	}
}

func (s *Server) RegisterRoutes() {

	s.router.Get("/healthz", handler.NewHealthCheckHandler().GetHandler())
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
