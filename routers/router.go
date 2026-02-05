package routers

import (
	"intern_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/v1/cars", controllers.AddCar)
	app.Get("/api/v1/cars", controllers.GetCars)
	app.Patch("/api/v1/cars/:id", controllers.UpdateCar)
	app.Delete("/api/v1/cars/:id", controllers.DeleteCar)

	// --- TAMBAHKAN BARIS INI ---
	// Mendaftarkan endpoint upload agar bisa dipanggil oleh frontend
	app.Post("/api/v1/upload", controllers.UploadFile)
	// ---------------------------
// --- TAMBAHAN BARU: ARCHIVE TOGGLE ---
    // Method PATCH karena kita hanya mengubah sebagian data (is_archived)
    app.Patch("/api/v1/cars/:id/archive", controllers.ToggleArchiveCar)
    // -------------------------------------
}