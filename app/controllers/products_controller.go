package controllers

import (
	"restful_auth/app/db"
	"restful_auth/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct{}

func (ProductsController) Index(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perPage"))
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	var countProducts int64
	db.DB.Model(&models.Product{}).Count(&countProducts)
	products := []models.Product{}
	if result := db.DB.Limit(perPage).Offset((page * perPage) - perPage).Order("id DESC").Find(&products); result.Error != nil {
		panic(result.Error.Error())
	}
	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total": countProducts,
		},
	})
}
