package utilities

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Checks if the token is authorized
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
