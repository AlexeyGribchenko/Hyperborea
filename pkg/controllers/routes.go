package controllers

import (
	"Marketplace/pkg/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB     *gorm.DB
	Config *config.Config
}

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	h := &handler{
		DB:     db,
		Config: cfg,
	}

	api := app.Group("/api/v1")

	productRoutes := api.Group("/products")
	productRoutes.Post("/", h.AddProduct)
	productRoutes.Get("/:id", h.GetProduct)
	//productRoutes.Get("/", h.GetAllProducts)
	productRoutes.Delete("/:id", h.DeleteProduct)

	productCharacteristicRoutes := productRoutes.Group("/:id/characteristics")
	productCharacteristicRoutes.Get("/", h.GetCharacteristics)
	productCharacteristicRoutes.Post("/", h.AddCharacteristic)
	productCharacteristicRoutes.Delete("/:index", h.DeleteCharacteristic)

	authRoutes := api.Group("/auth")
	authRoutes.Post("/sign-up", h.SignUp)
	authRoutes.Get("/sign-in", h.SignIn)

	cartRoutes := api.Group("/cart")
	cartRoutes.Get("/", h.GetProductsInCart)
	cartRoutes.Post("/", h.AddProductToCart)
	cartRoutes.Delete("/", h.DeleteProductFromCart)

	favouritesRoutes := api.Group("/favourites")
	favouritesRoutes.Get("/", h.GetProductsInFavourites)
	favouritesRoutes.Post("/", h.AddProductToFavourites)
	favouritesRoutes.Delete("/", h.DeleteProductFromFavourites)

	/*orderRoutes := api.Group("/order")
	orderRoutes.Get("/", h.GetAllOrders)
	orderRoutes.Get("/:id", h.GetOrder)
	orderRoutes.Post("/", h.CreateOrder)

	shopRoutes := api.Group("/shop")
	shopRoutes.Post("/", h.CreateShop)
	shopRoutes.Get("/", h.GetShop)

	commentRoutes := productRoutes.Group("/:id/comment")
	commentRoutes.Get("/", h.GetAllComments)
	commentRoutes.Post("/", h.CreateComment)*/

}
