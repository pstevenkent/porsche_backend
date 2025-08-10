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

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	config.ConnectDatabase()

	defer config.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	routers.SetUp(app)

	log.Fatal(app.Listen(":8080"))
	
}