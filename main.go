// Create server with Fiber
package main

import (
	"github.com/joho/godotenv"
	"github.com/karlbehrensg/go-clean-architecture-template/bootstrap"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	server := bootstrap.NewServer()
	server.Start()
}
