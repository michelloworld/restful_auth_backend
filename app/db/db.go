package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBError error

func Init() {
	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	DB = db
	DBError = err
}
