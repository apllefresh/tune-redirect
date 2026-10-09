package api

import (
	"os"

	"github.com/apllefresh/tune-redirect/internal/handler"
)

type Container struct {
	server *Server
}

func NewContainer() *Container {
	httpAddress := os.Getenv("HTTP_ADDR")
	if httpAddress == "" {
		panic("HTTP_ADDR environment variable not set")
	}

	server := NewServer(httpAddress)

	healthCheckHandler := handler.NewHealthCheckHandler()
	err := server.AddRoute(healthCheckHandler)
	if err != nil {
		panic(err)
	}

	return &Container{
		server: server,
	}
}

func (s *Container) Start() error {
	return s.server.httpServer.ListenAndServe()
}
