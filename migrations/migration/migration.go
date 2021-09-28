package main

import (
	"fmt"
	"log"
	"restful_auth/app/db"
	"restful_auth/app/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()

	db.DB.AutoMigrate(&models.User{}, &models.Product{})
	fmt.Println("Complete Migration!!!")
}
