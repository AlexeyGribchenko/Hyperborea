package controllers

import (
	"Marketplace/pkg/models"
	"Marketplace/pkg/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AddProductToFavouritesRequest struct {
	ProductId int `json:"productId"`
}

func (h handler) GetProductsInFavourites(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	userId, err := services.GetUserId(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var productsInFavourites []models.Product

	if result := h.DB.Debug().Table("t_products").
		Select("t_products.*").
		Joins("INNER JOIN t_product_to_favourites ON t_products.id = t_product_to_favourites.product_id").
		Where("t_product_to_favourites.user_id = ?", userId).
		Scan(&productsInFavourites); result != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(&productsInFavourites)
}

func (h handler) AddProductToFavourites(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")

	userId, err := services.GetUserId(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var body AddProductToCartRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var productToFavourites models.ProductToFavourite

	productToFavourites.ProductId = body.ProductId
	productToFavourites.UserId = userId
	productToFavourites.AddDate = time.Now()

	if result := h.DB.Create(&productToFavourites); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h handler) DeleteProductFromFavourites(c *fiber.Ctx) error {
	id := c.Params("id")

	if result := h.DB.
		Where("id = ?", id).
		Delete(&models.ProductToFavourite{}); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
