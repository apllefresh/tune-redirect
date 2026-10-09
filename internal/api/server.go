package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func NewServer(httpAddress string) *Server {
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

type IHandler interface {
	GetHandler() http.HandlerFunc
	GetRoute() string
	GetMethod() string
}

func (s *Server) AddRoute(handler IHandler) error {
	switch handler.GetMethod() {
	case http.MethodGet:
		s.router.Get(handler.GetRoute(), handler.GetHandler())
	default:
		return errors.New("invalid method")
	}

	return nil
}
