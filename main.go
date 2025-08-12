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
	app := fiber.New(fiber.Config{
		BodyLimit: 250 * 1024 * 1024, // 250 MB
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	config.ConnectDatabase()
	defer config.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))

	// --- TAMBAHKAN BARIS INI ---
	// Baris ini akan membuat semua file di dalam folder `./uploads`
	// dapat diakses melalui URL, contoh: http://localhost:8080/uploads/namafile.jpg
	app.Static("/uploads", "./uploads")
	// ---------------------------

	routers.SetUp(app)

	log.Fatal(app.Listen(":8080"))
}