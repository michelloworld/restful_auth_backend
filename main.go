package main

import (
	"log"
	"restful_auth/app/db"
	"restful_auth/app/models"
	"restful_auth/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))

	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(fiber.Map{"error": fiber.Map{"message": r}})
			}
		}()
		return c.Next()
	})

	db.Init()
	db.DB.Migrator().CreateTable(&models.User{})
	db.DB.Migrator().CreateTable(&models.Product{})

	routes.Init(app)

	app.Listen(":4001")
}
