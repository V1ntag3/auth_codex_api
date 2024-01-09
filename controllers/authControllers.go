package controllers

import (
	"auth_codex_api/database"
	"auth_codex_api/html"
	"auth_codex_api/models"
	"auth_codex_api/utilities"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register user
func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// validation data
	error_validation := make(map[string]string)

	if utilities.ContainsNumber(data["name"]) || utilities.OnlyEmptySpaces(data["name"]) {
		error_validation["name"] = "Name is invalid"
	}

	if utilities.ContainsNumber(data["surname"]) || utilities.OnlyEmptySpaces(data["surname"]) {
		error_validation["surname"] = "Surname is invalid"
	}

	if !utilities.IsValidEmail(data["email"]) {
		error_validation["email"] = "E-mail is invalid"
	}

	if !utilities.IsValidPassword(data["password"]) {
		error_validation["password"] = "Password is invalid"
	}

	if data["app"] == "" {
		error_validation["app"] = "App is invalid"
	}

	if len(error_validation) != 0 {
		return c.Status(400).JSON(error_validation)
	}

	// hash of password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// format number
	replacer := strings.NewReplacer("(", "", ")", "", "-", "", " ", "")
	mobile := replacer.Replace(data["mobile"])

	user := models.User{
		Id:        uuid.New().String(),
		Name:      data["name"],
		Surname:   data["surname"],
		Email:     strings.ToLower(data["email"]),
		Mobile:    strings.TrimSpace(mobile),
		Password:  password,
		CreatedAt: utilities.DateTimeNow(),
		Apps: []models.App{
			{Id: data["app"]},
		},
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			error_validation["email"] = "E-mail already registered"
			return c.Status(400).JSON(error_validation)
		}
		if err.Error() == "UNIQUE constraint failed: users.mobile" {
			error_validation["mobile"] = "Mobile already registered"
			return c.Status(400).JSON(error_validation)
		}
	}
	html.HtmlEmailVerify(data["email"], user)

	return c.JSON(user)
}

// Login
func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.ErrInternalServerError
	}

	var user models.User

	if err := database.DB.Where("email= ?", data["email"]).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.Id,
		ExpiresAt: jwt.NewNumericDate(utilities.DateTimeNowAddHours(24)),
	})

	token, err := claims.SignedString([]byte(utilities.GoDotEnvVariable("SECRETKEY")))
	if err != nil {
		print(err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.Status(404).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	utilities.AuthorizedToken(token)

	return c.JSON(fiber.Map{
		"token":   token,
		"expires": utilities.DateTimeNowAddHours(24),
	})
}

// Logs the user out
func Logout(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	utilities.UnauthorizedToken(token.Raw)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
