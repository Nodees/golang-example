package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GetUserFromToken(token string, secret string) (jwt.MapClaims, error) {

	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, _ := tokenByte.Claims.(jwt.MapClaims)

	return claims, nil
}

func In(value any, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
