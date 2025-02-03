// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"trailer_chatbot/handlers"
)

// SetupRouter initializes the Fiber app and routes
func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Post("/chat", handlers.ChatbotHandler)
	return app
}

