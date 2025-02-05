// handlers/chatbot.go
package handlers

import (
	"fmt"
	"strings"
	"trailer_chatbot/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var activeConnections = make(map[*websocket.Conn]bool)

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

// DashboardHandler provides real-time dashboard statistics
func DashboardHandler(c *fiber.Ctx) error {
	var totalTrailers, completed, pending int64
	database.DB.Model(&database.Trailer{}).Count(&totalTrailers)
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Completed").Count(&completed)
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Pending").Count(&pending)

	return c.JSON(fiber.Map{
		"total_trailers": totalTrailers,
		"completed":      completed,
		"pending":        pending,
	})
}

// processChat processes chatbot messages
func processChat(message string) string {
	switch {
	case message == "list trailers":
		return getTrailerList()
	case strings.HasPrefix(message, "assign trailer"):
		return assignTrailer(message)
	case strings.HasPrefix(message, "status of trailer"):
		return getTrailerStatus(message)
	case message == "generate unloading report":
		return generateReport()
	case strings.HasPrefix(message, "find trailer"):
		return findTrailer(message)
	case message == "notify unloading complete":
		return notifyUnloadingComplete()
	case message == "dashboard status":
		return getDashboardStatus()
	default:
		return "I can help with unloading tasks. Try 'list trailers', 'assign trailer <ID> to dock <DockName>', 'status of trailer <ID>', 'generate unloading report', or 'find trailer <ID>'."
	}
}

// notifyUnloadingComplete sends a message to all active WebSocket connections
func notifyUnloadingComplete() string {
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Assigned").Update("status", "Completed")
	for conn := range activeConnections {
		conn.WriteMessage(websocket.TextMessage, []byte("A trailer has finished unloading!"))
	}
	return "All assigned trailers marked as completed, and notifications sent."
}

// findTrailer retrieves trailer details by ID
func findTrailer(message string) string {
	parts := strings.Split(message, " ")
	if len(parts) < 3 {
		return "Invalid format. Try 'find trailer <ID>'"
	}
	trailerID := parts[2]

	var trailer database.Trailer
	database.DB.Where("id = ?", trailerID).First(&trailer)

	if trailer.ID == 0 {
		return "Trailer not found."
	}
	return fmt.Sprintf("Trailer %d is at dock %s with status: %s.", trailer.ID, trailer.DockingBay, trailer.Status)
}

// generateReport provides unloading statistics
func generateReport() string {
	var completed, pending int64
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Completed").Count(&completed)
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Pending").Count(&pending)

	return fmt.Sprintf("Unloading Report:\n- Completed Trailers: %d\n- Pending Trailers: %d", completed, pending)
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

// assignTrailer assigns a trailer to a specific dock
func assignTrailer(message string) string {
	parts := strings.Split(message, " ")
	if len(parts) < 6 {
		return "Invalid format. Try 'assign trailer <ID> to dock <DockName>'"
	}
	trailerID := parts[2]
	dockName := parts[5]

	// Update database
	database.DB.Model(&database.Trailer{}).Where("id = ?", trailerID).Update("dock", dockName)

	return fmt.Sprintf("Trailer %s assigned to Dock %s.", trailerID, dockName)
}

// getTrailerStatus retrieves the unloading status of a specific trailer
func getTrailerStatus(message string) string {
	parts := strings.Split(message, " ")
	if len(parts) < 4 {
		return "Invalid format. Try 'status of trailer <ID>'"
	}
	trailerID := parts[3]

	var trailer database.Trailer
	database.DB.Where("id = ?", trailerID).First(&trailer)

	if trailer.ID == 0 {
		return "Trailer not found."
	}
	return fmt.Sprintf("Trailer %s is currently %s.", trailerID, trailer.Status)
}

// getDashboardStatus fetches trailer statistics for dashboard
func getDashboardStatus() string {
	var totalTrailers, completed, pending int64
	database.DB.Model(&database.Trailer{}).Count(&totalTrailers)
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Completed").Count(&completed)
	database.DB.Model(&database.Trailer{}).Where("status = ?", "Pending").Count(&pending)

	return fmt.Sprintf("Dashboard Status:\n- Total Trailers: %d\n- Completed: %d\n- Pending: %d", totalTrailers, completed, pending)
}
