// main.go
package main

import (
	"log"
	"trailer_chatbot/database"
	"trailer_chatbot/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.InitializeDatabase()

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Allow requests from frontend
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Register routes
	routes.SetupRouter(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
