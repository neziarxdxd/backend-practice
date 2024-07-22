package main

import "backend-practice/routes"

func main() {
	// Load the routes
	r := routes.SetupRouter()

	// Initialize database
	// models.SetupDatabase()

	// Start the HTTP API
	r.Run()
	println("Hello, World!")
}
