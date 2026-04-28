// Package main is the entry point for the Porsche Catalog Backend API.
// It initializes the Fiber web server, connects to MongoDB, configures
// CORS middleware, and registers all API routes.
package main

import (
	"intern_backend/config"
	"intern_backend/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load environment variables from .env file (MONGO_URI, DB_NAME, Cloudinary credentials)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	// Establish MongoDB connection and ensure it closes on shutdown
	config.ConnectDatabase()
	defer config.DisconnectDatabase()

	// Enable CORS for all origins so the frontend (localhost:5173) can reach the API
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Register all API route handlers
	routers.SetUp(app)

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
