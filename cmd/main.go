package main

import (
	"Marketplace/pkg/config"
	"Marketplace/pkg/controllers"
	"Marketplace/pkg/db"
	"Marketplace/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	app.Use(middlewares.JWTMiddleware())
	controllers.RegisterRoutes(app, h, &c)

	if err := app.Listen(c.Port); err != nil {
		return
	}
}
