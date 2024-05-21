package controllers

import (
	"Marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AddCharacteristicRequest struct {
	ProductId int    `json:"productId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

func (h handler) AddCharacteristic(c *fiber.Ctx) error {

	body := AddCharacteristicRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var characteristic models.ProductCharacteristic

	characteristic.ProductId = productId
	characteristic.Name = body.Name
	characteristic.Value = body.Value

	h.DB.Create(&characteristic)

	return c.Status(fiber.StatusCreated).JSON(&characteristic)
}

func (h handler) GetCharacteristics(c *fiber.Ctx) error {
	productId := c.Params("id")

	var characteristics []models.ProductCharacteristic

	if result := h.DB.Debug().Table("t_product_characteristic pc").
		Select("pc.*").
		Joins("INNER JOIN t_product p ON p.id = pc.product_id").
		Where("p.id = ?", productId).
		Scan(&characteristics); result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(&characteristics)
}

func (h handler) DeleteCharacteristic(c *fiber.Ctx) error {
	characteristicId := c.Params("index")

	if result := h.DB.
		Where("id = ?", characteristicId).
		Delete(&models.ProductCharacteristic{}); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("Deleted")
}
