package utilities

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Verifica se o token está autorizado
func IsAuthenticadToken(c *fiber.Ctx, SecretKey string) (*jwt.Token, error) {

	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) <= 1 {
		return nil, errors.New("unauthorized")
	}
	reqToken = splitToken[1]

	if !IsAuthorizedToken(reqToken) {
		return nil, errors.New("unauthorized")
	}

	token, err := jwt.ParseWithClaims(reqToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, errors.New("unauthorized")
	}

	return token, nil
}
