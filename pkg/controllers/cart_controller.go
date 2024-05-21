package controllers

import (
	"Marketplace/pkg/models"
	"Marketplace/pkg/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AddProductToCartRequest struct {
	ProductId int `json:"productId"`
	Amount    int `json:"amount"`
}

type UpdateProductAmountRequest struct {
	NewAmount int `json:"newAmount"`
}

type ProductInCartResponse struct {
	SmallProduct struct {
		Id     int     `json:"id"`
		Title  string  `json:"title"`
		Amount string  `json:"amount"`
		Price  float32 `json:"price"`
	} `json:"product" gorm:"embedded"`
	Amount int `json:"amount"`
}

func (h handler) GetProductsInCart(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	userId, err := services.GetUserId(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var productsInCart []ProductInCartResponse

	if result := h.DB.Debug().Table("t_product p").
		Select("p.id, p.title, pc.amount, pp.price").
		Joins("INNER JOIN t_product_to_cart pc ON p.id = pc.product_id").
		Joins("INNER JOIN t_product_price pp ON pp.id = p.current_price_id").
		Where("pc.user_id = ?", userId).
		Scan(&productsInCart); result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(&productsInCart)
}

func (h handler) AddProductToCart(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")

	userId, err := services.GetUserId(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var body AddProductToCartRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var productToCart models.ProductToCart

	productToCart.ProductId = body.ProductId
	productToCart.UserId = userId
	productToCart.AddDate = time.Now()
	productToCart.Amount = body.Amount

	if result := h.DB.Create(&productToCart); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h handler) DeleteProductFromCart(c *fiber.Ctx) error {
	id := c.Params("id")

	if result := h.DB.Delete(id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h handler) ChangeProductAmount(c *fiber.Ctx) error {
	id := c.Params("id")

	var body UpdateProductAmountRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if result := h.DB.Find(id).Update("amount", body.NewAmount); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
