package controllers

import (
	"intern_backend/constants"
	"intern_backend/controllers/helper"
	"intern_backend/models"
	"intern_backend/output"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddCar sekarang dimodifikasi untuk menangani multipart/form-data (termasuk file)
func AddCar(c *fiber.Ctx) error {
	// 1. Ambil form multipart dari permintaan. Ini bisa menangani teks dan file.
	form, err := c.MultipartForm()
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, "Gagal mem-parsing form: "+err.Error())
	}

	// 2. Ambil semua data teks dari form.
	vehicle := form.Value["vehicle"][0]
	modelYearStr := form.Value["modelyear"][0]
	exteriorColour := form.Value["exteriorcolour"][0]
	interiorColours := form.Value["interiorcolours"][0]
	wheels := form.Value["wheels"][0]
	seats := form.Value["seats"][0]
	roofTransport := form.Value["rooftransport"][0]
	infotainment := form.Value["infotainment"][0]
	powertrainPerformance := form.Value["powertrainperformance"]

	// 3. Ambil dan simpan file gambar.
	files := form.File["images"]
	var imageUrls []string

	// Buat folder 'uploads' di direktori backend jika belum ada.
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, "Gagal membuat direktori uploads: "+err.Error())
	}

	for _, file := range files {
		// Simpan setiap file yang diunggah ke folder lokal.
		filePath := "./uploads/" + file.Filename
		if err := c.SaveFile(file, filePath); err != nil {
			return output.GetError(c, fiber.StatusInternalServerError, "Gagal menyimpan file: "+err.Error())
		}
		// Simpan path file untuk dimasukkan ke database.
		imageUrls = append(imageUrls, filePath)
	}

	// 4. Buat objek Car baru dengan data yang sudah diproses.
	modelYear, _ := strconv.Atoi(modelYearStr)
	car := models.Car{
		Vehicle:               vehicle,
		ModelYear:             modelYear,
		Images:                imageUrls,
		ExteriorColour:        exteriorColour,
		InteriorColours:       interiorColours,
		Wheels:                wheels,
		Seats:                 seats,
		RoofTransport:         roofTransport,
		Infotainment:          infotainment,
		PowertrainPerformance: powertrainPerformance,
	}

	// 5. Masukkan objek Car ke database.
	res, err := helper.InsertData(string(constants.Cars), &car)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Kirim respons sukses.
	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})
}

// Fungsi-fungsi di bawah ini tidak diubah dan akan tetap berfungsi.
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