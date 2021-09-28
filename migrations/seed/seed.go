package main

import (
	"fmt"
	"log"
	"math/rand"
	"restful_auth/app/db"
	"restful_auth/app/models"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	email := "admin@admin.com"
	findUser := models.User{}
	result := db.DB.Where("email = ?", email).Find(&findUser)
	if result.RowsAffected == int64(0) {
		user := models.User{Email: email, Password: "password"}
		db.DB.Create(&user)
	}

	findProducts := models.Product{}
	result1 := db.DB.Find(&findProducts)
	if result1.RowsAffected == int64(0) {
		for i := 0; i < 100; i++ {
			product := models.Product{Name: "Product " + strconv.Itoa(i), Price: float32(rand.Intn(100)) * rand.Float32()}
			db.DB.Create(&product)
		}
	}

	fmt.Println("Complete Seed!!!")
}
