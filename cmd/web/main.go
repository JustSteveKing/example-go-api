package main

import (
	"github.com/JustSteveKing/example-go-api/pkg/kernel"

	"github.com/joho/godotenv"
)

func main() {
	// Load our environment variables
	if err := godotenv.Load(); err != nil {
		panic("No .env found")
	}

	// Create our application
	app := kernel.Boot()

	// Load our Routes

	// Run our server
	go func() {
		app.Run()
	}()

	// wait for shutdown signal
	app.WaitForShutdown()
}
