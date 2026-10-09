package main

import (
	"os"

	"github.com/apllefresh/tune-redirect/internal/api"
)

func main() {
	httpAddress := os.Getenv("HTTP_ADDR")
	if httpAddress == "" {
		panic("HTTP_ADDR environment variable not set")
	}

	server := api.NewServer(httpAddress)
	server.RegisterRoutes()

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
