// handlers/chatbot.go
package handlers

import (
	"fmt"
	"trailer_chatbot/database"

	"github.com/gofiber/fiber/v2"
)

// ChatbotHandler handles chatbot requests
func ChatbotHandler(c *fiber.Ctx) error {
	type Request struct {
		Message string `json:"message"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	response := processChat(req.Message)
	return c.JSON(fiber.Map{"response": response})
}

// processChat processes chatbot messages
func processChat(message string) string {
	switch message {
	case "status":
		return getTrailerStatus()
	case "assign":
		return assignTrailer()
	case "complete":
		return completeTrailerTasks()
	case "list trailers":
		return getTrailerList()
	default:
		return "I can help with unloading tasks. Try 'list trailers', 'status', 'assign', or 'complete'."
	}
}

// getTrailerStatus retrieves the status of trailers
func getTrailerStatus() string {
	var trailers []database.Trailer
	database.DB.Find(&trailers)
	if len(trailers) == 0 {
		return "No trailers currently in process."
	}
	var statusResponse string
	for _, trailer := range trailers {
		statusResponse += fmt.Sprintf("Trailer %s at Dock %s: %s\n", trailer.Number, trailer.DockingBay, trailer.Status)
	}
	return statusResponse
}

// assignTrailer assigns a new trailer
func assignTrailer() string {
	newTrailer := database.Trailer{Number: "TR123", DockingBay: "5A", Status: "Assigned"}
	database.DB.Create(&newTrailer)
	return fmt.Sprintf("Trailer %s assigned to Dock %s.", newTrailer.Number, newTrailer.DockingBay)
}

// completeTrailerTasks marks assigned trailers as completed
func completeTrailerTasks() string {
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Assigned").Update("status", "Completed")
	return "All assigned trailers marked as completed."
}

// getTrailerList retrieves the list of trailers from the database
func getTrailerList() string {
	var trailers []database.Trailer
	database.DB.Find(&trailers)

	if len(trailers) == 0 {
		return "No trailers found."
	}

	var trailerList string
	for i, trailer := range trailers {
		trailerList += fmt.Sprintf("%d. %d\n", i+1, trailer.ID)
	}

	return "Here are the available trailers:\n" + trailerList
}
