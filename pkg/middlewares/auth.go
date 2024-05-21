package middlewares

import (
	"Marketplace/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" || strings.HasPrefix(tokenString, "Bearer ") {
			return c.Next()
		}
		tokenString, _ = strings.CutPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, services.TokenFunc)

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		return c.Next()
	}
}

//
//func main() {
//	app := fiber.New()
//
//	// Middleware для логирования запросов
//	app.Use(logger.New())
//
//	// Middleware для обработки CORS
//	app.Use(cors.New())
//
//	// Использование нашего JWTMiddleware
//	app.Use(JWTMiddleware())
//
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString("Привет, мир!")
//	})
//
//	log.Fatal(app.Listen(":3000"))
//}
//Обратите внимание, что в этом примере middleware JWTMiddleware() проверяет наличие JWT токена в заголовке Authorization, затем проверяет его валидность, используя секретный ключ (в реальном приложении вам нужно будет использовать ваш секретный ключ для верификации токена).
//
//Пожалуйста, не забудьте заменить "YOUR_SECRET_KEY" на ваш реальный секретный ключ перед использованием в производственном приложении.🧐🔒
