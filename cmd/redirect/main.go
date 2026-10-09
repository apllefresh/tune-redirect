package main

import (
	"github.com/apllefresh/tune-redirect/internal/http"
)

func main() {
	server := http.NewServer()
	server.RegisterRoutes()

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
