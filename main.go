package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	// WebSocket route (remove HEAD route if not needed)
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			msgType, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			if err := c.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	}))

	// Log routes to check for conflicts
	for _, route := range app.GetRoutes() {
		log.Println("Route:", route)
	}

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
