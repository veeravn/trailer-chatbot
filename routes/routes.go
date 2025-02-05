// routes/routes.go
package routes

import (
	"trailer_chatbot/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Post("/chat", handlers.ChatbotHandler)
	app.Post("/dashboard", handlers.DashboardHandler)
}
