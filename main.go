// main.go
package main

import (
	"log"
	"trailer_chatbot/database"
	"trailer_chatbot/routes"
)

func main() {
	database.InitializeDatabase()
	app := routes.SetupRouter()
	log.Fatal(app.Listen(":3000"))
}

