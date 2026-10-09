package main

import (
	"github.com/apllefresh/tune-redirect/internal/api"
)

func main() {
	var container = api.NewContainer()

	err := container.Start()
	if err != nil {
		panic(err)
	}
}
