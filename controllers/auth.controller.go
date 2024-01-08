package controllers

import (
	"core/configs"
	connection "core/connections"
	"core/models"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(user models.User) (string, *jwt.Token, error) {
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	conf, err := configs.LoadConfig(".")
	if err != nil {
		return "", nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	claims["sub"] = user.ID
	claims["exp"] = now.Add(conf.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return "", nil, fmt.Errorf("generating JWT token failed: %v", err)
	}

	return tokenString, tokenByte, nil
}

func LoginHandler(c *fiber.Ctx) error {
	var payload *models.User

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	conf, erro := configs.LoadConfig(".")
	if erro != nil {
		log.Fatal("Não foi possivel carregar variaveis de ambiente: ", erro)
	}

	var user models.User
	result := connection.DB.First(&user, "tx_username = ?", payload.Username)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid username or Password"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid username or Password"})
	}

	fmt.Println("token", c.Cookies("token"))

	if c.Cookies("token") != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": c.Cookies("token")})
	}
	tokenString, _, err := GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:   "token",
		Value:  tokenString,
		Path:   "/",
		MaxAge: int(conf.JwtMaxAge * 60),
		Secure: false,
		Domain: conf.Domain,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
