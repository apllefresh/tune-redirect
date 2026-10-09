package main

import (
	"os"

	"github.com/apllefresh/tune-redirect/internal/http"
)

func main() {
	httpAddress := os.Getenv("HTTP_ADDR")

	server := http.NewServer(httpAddress)
	server.RegisterRoutes()

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
