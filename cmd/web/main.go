package main

import (
	"github.com/JustSteveKing/example-go-api/pkg/kernel"
	"github.com/JustSteveKing/example-go-api/pkg/routes"

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
	routes.Load(app)

	// Run our server
	go func() {
		app.Run()
	}()

	// wait for shutdown signal
	app.WaitForShutdown()
}
