package routes

import (
	"fmt"
	"os"
	"restful_auth/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Init(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/auth/token", controllers.AuthController{}.Token)
	api.Post("/auth/refresh_token", controllers.AuthController{}.RefreshToken)
	api.Post("/auth/sign_out", controllers.AuthController{}.SignOut)

	api.Use(func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")
		tokenString := bearerToken[7:]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
		})

		if err != nil {
			c.Status(401)
			return c.JSON(fiber.Map{
				"error": fiber.Map{
					"message": err.Error(),
				},
			})
		}

		return c.Next()
	})

	api.Get("/products", controllers.ProductsController{}.Index)
}
