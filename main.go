// main.go
package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"trailer_chatbot/database"
	"trailer_chatbot/routes"
)

func main() {
	database.InitializeDatabase()
	app := routes.SetupRouter()
	log.Fatal(app.Listen(":3000"))
}

