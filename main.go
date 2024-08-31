package main

import (
	"log"
	"os"

	"github.com/amanasmuei/gofiber-fcm.git/gofiberfcm"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Firebase
	if err := gofiberfcm.Init(); err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	// Create a new Fiber application
	app := fiber.New()

	// Middleware for logging
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	// Route to send push notification
	app.Post("/send-notification", func(c *fiber.Ctx) error {
		var req struct {
			Token string `json:"token"`
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		// Parse request body
		if err := c.BodyParser(&req); err != nil {
			log.Printf("Error parsing request body: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		// Send notification
		messageID, err := gofiberfcm.SendNotification(req.Token, req.Title, req.Body)
		if err != nil {
			log.Printf("Error sending notification: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to send notification",
			})
		}

		return c.JSON(fiber.Map{
			"message":    "Notification sent successfully",
			"message_id": messageID,
		})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port
	}
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
