package main

import (
	"github.com/joho/godotenv"
)

func main() {
	// Load our environment variables
	if err := godotenv.Load(); err != nil {
		panic("No .env found")
	}
}
