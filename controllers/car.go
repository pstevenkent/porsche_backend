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
	res, err := helper.InsertData(string(constants.Cars), &car)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}
	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})
}

// Fungsi lain tidak perlu diubah
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