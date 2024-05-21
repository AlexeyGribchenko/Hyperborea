package controllers

import (
	"Marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AddProductRequestBody struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ShopId        int     `json:"shopId"`
	Price         float32 `json:"price"`
	PriorityValue int     `json:"priorityValue"`
}

func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product
	var productPrice models.ProductPrice

	product.Title = body.Title
	product.Description = body.Description
	product.CurrentPriceId = -1
	product.PriorityValue = body.PriorityValue
	product.ShopId = body.ShopId

	if result := h.DB.Create(&product); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	productPrice.Price = body.Price
	productPrice.ProductId = product.Id
	productPrice.Discount = 0
	productPrice.CreationDate = time.Now()

	if result := h.DB.Create(&productPrice); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	product.CurrentPriceId = productPrice.Id

	if result := h.DB.Save(&product); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}

func (h handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h handler) GetProducts(c *fiber.Ctx) error {

	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (h handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if result := h.DB.Delete(id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("Deleted")
}
