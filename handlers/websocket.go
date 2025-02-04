// handlers/websocket.go
package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

// WebSocketHandler handles real-time chat via WebSockets
func WebSocketHandler(c *websocket.Conn) {
	defer c.Close()

	for {
		// Read message from client
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Echo the message back to the client
		err = c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}
