package controllers

import (
	"intern_backend/constants"
	"intern_backend/controllers/helper"
	"intern_backend/models"
	"intern_backend/output"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddCar sekarang diperbarui untuk menerima CommNr, Price, dan file
func AddCar(c *fiber.Ctx) error {
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	// Default IsArchived false saat create
	car.IsArchived = false 
	
	res, err := helper.InsertData(string(constants.Cars), &car)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}
	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})
}

func UpdateCar(c *fiber.Ctx) error {
	var car models.Car
	if err := c.BodyParser(&car); err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	_, err = helper.UpdateData(string(constants.Cars), "_id", objId, &car)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	return output.GetSuccess(c, string(constants.SuccessUpdateMessage), fiber.Map{})
}

func GetCars(c *fiber.Ctx) error {
	var cars []models.Car
	_, err := helper.RetrieveData(bson.M{}, string(constants.Cars), &cars)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}
	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{"cars": cars})
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	_, err = helper.DeleteData(string(constants.Cars), "_id", objId)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}
	return output.GetSuccess(c, string(constants.SuccessDeleteMessage), fiber.Map{})
}

// --- FUNGSI BARU DI BAWAH INI ---

// ToggleArchiveCar mengubah status is_archived (true <-> false)
func ToggleArchiveCar(c *fiber.Ctx) error {
	// 1. Ambil ID dari Parameter URL
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	// 2. Ambil Data Mobil Saat Ini (untuk tahu status true/false nya sekarang)
	var cars []models.Car
	// Menggunakan helper.RetrieveData dengan filter _id
	_, err = helper.RetrieveData(bson.M{"_id": objId}, string(constants.Cars), &cars)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}
	
	// Cek apakah mobil ditemukan
	if len(cars) == 0 {
		return output.GetError(c, fiber.StatusNotFound, "Car not found")
	}

	// 3. Lakukan Toggle (Balik logika)
	carToUpdate := cars[0]
	newStatus := !carToUpdate.IsArchived
	
	// Kita buat map update khusus agar tidak menimpa field lain secara tidak sengaja,
	// TAPI karena helper.UpdateData Anda sepertinya menerima struct/interface, 
	// kita update struct yang sudah diambil tadi.
	carToUpdate.IsArchived = newStatus

	// 4. Simpan Perubahan ke Database
	// Menggunakan helper.UpdateData yang sudah ada
	_, err = helper.UpdateData(string(constants.Cars), "_id", objId, &carToUpdate)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, "Car archive status updated", fiber.Map{
		"is_archived": newStatus,
		"car_id":      id,
	})
}