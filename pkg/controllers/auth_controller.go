package controllers

import (
	"Marketplace/pkg/models"
	"Marketplace/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type SignUpRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtAccessTokenResponse struct {
	Token string `json:"token"`
}

func (h handler) SignUp(c *fiber.Ctx) error {

	body := SignUpRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Password = body.Password
	user.Email = body.Email
	user.PhoneNumber = body.PhoneNumber
	user.Name = body.Name

	tokenString, err := services.GenerateAccessToken(&user, h.Config.SecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	response := JwtAccessTokenResponse{Token: tokenString}

	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h handler) SignIn(c *fiber.Ctx) error {

	body := SignInRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if user.Password != body.Password {
		return c.Status(fiber.StatusForbidden).JSON("Bad credentials")
	}

	tokenString, err := services.GenerateAccessToken(&user, h.Config.SecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	response := JwtAccessTokenResponse{Token: tokenString}

	return c.Status(fiber.StatusCreated).JSON(&response)
}
