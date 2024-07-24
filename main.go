package main

import (
	"backend-practice/models"
	"backend-practice/routes"
)

func main() {
	// Load the routes
	r := routes.SetupRouter()
	models.SetupDatabase()

	// Initialize database
	// models.SetupDatabase()

	// Start the HTTP API

	r.Run()
	println("Hello, World!")
}
