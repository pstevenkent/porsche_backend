package routers

import (
	"intern_backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/v1/cars", controllers.AddCar)

	app.Get("/api/v1/cars", controllers.GetCars)

	app.Put("/api/v1/cars/:id", controllers.UpdateCar)

	app.Delete("/api/v1/cars/:id", controllers.DeleteCar)

}
