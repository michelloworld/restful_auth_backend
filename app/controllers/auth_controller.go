package controllers

import (
	"database/sql"
	"math/rand"
	"os"
	"restful_auth/app/db"
	"restful_auth/app/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func generateAccessToken(email string) (string, error) {
	tokenLife, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * time.Duration(int64(tokenLife))).Unix(),
		"nonce": time.Now().UnixNano() + int64(rand.Intn(10000)),
		// "nonce": int64(rand.Intn(100)),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	return tokenString, err
}

func generateRefreshToken(email string) (string, error) {
	refreshTokenLife, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE"))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * time.Duration(int64(refreshTokenLife))).Unix(),
		"nonce": time.Now().UnixNano() + int64(rand.Intn(10000)),
		// "nonce": int64(rand.Intn(100)),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	return refreshTokenString, err
}

func (AuthController) Token(c *fiber.Ctx) error {
	body := struct {
		Email    string
		Password string
	}{}
	if err := c.BodyParser(&body); err != nil {
		panic(err.Error())
	}

	user := models.User{}
	if result := db.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		panic(result.Error.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		panic(err.Error())
	}

	tokenString, err := generateAccessToken(body.Email)
	if err != nil {
		panic(err.Error())
	}

	refreshTokenString, err := generateAccessToken(body.Email)
	if err != nil {
		panic(err.Error())
	}

	refreshTokenLife, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE"))
	refreshTokenExpiresIn := time.Now().Add(time.Second * time.Duration(int64(refreshTokenLife)))
	if result := db.DB.Model(&user).Updates(models.User{RefreshToken: sql.NullString{String: refreshTokenString, Valid: true}, ExpiresIn: sql.NullTime{Time: refreshTokenExpiresIn, Valid: true}}); result.Error != nil {
		panic(result.Error.Error())
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"email":        user.Email,
			"accessToken":  tokenString,
			"refreshToken": refreshTokenString,
		},
	})
}

func (AuthController) RefreshToken(c *fiber.Ctx) error {
	body := struct {
		RefreshToken string
	}{}
	if err := c.BodyParser(&body); err != nil {
		panic(err.Error())
	}

	findUser := models.User{}
	if result := db.DB.Where("refresh_token = ?", body.RefreshToken).First(&findUser); result.Error != nil {
		panic(result.Error.Error())
	}

	if time.Now().After(findUser.ExpiresIn.Time) {
		panic("refresh token was expires")
	}

	tokenString, err := generateAccessToken(findUser.Email)
	if err != nil {
		panic(err.Error())
	}

	refreshTokenString, err := generateAccessToken(findUser.Email)
	if err != nil {
		panic(err.Error())
	}

	refreshTokenLife, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE"))
	expiresIn := time.Now().Add(time.Second * time.Duration(int64(refreshTokenLife)))
	if result := db.DB.Model(&findUser).Updates(models.User{RefreshToken: sql.NullString{String: refreshTokenString, Valid: true}, ExpiresIn: sql.NullTime{Time: expiresIn, Valid: true}}); result.Error != nil {
		panic(result.Error.Error())
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"accessToken":  tokenString,
			"refreshToken": refreshTokenString,
		},
	})
}

func (AuthController) SignOut(c *fiber.Ctx) error {
	body := struct {
		RefreshToken string
	}{}
	if err := c.BodyParser(&body); err != nil {
		panic(err.Error())
	}

	findUser := models.User{}
	if result := db.DB.Where("refresh_token = ? ", body.RefreshToken).First(&findUser); result.Error != nil {
		panic(result.Error.Error())
	}

	if result := db.DB.Model(&findUser).Updates(models.User{RefreshToken: sql.NullString{String: "null", Valid: false}, ExpiresIn: sql.NullTime{Time: time.Now(), Valid: false}}); result.Error != nil {
		panic(result.Error.Error())
	}

	c.SendStatus(204)
	return c.JSON(nil)
}
