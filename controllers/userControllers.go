package controllers

import (
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// User data
func Profile(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.MapClaims)

	user_id, _ := claims.GetIssuer()

	var user models.User

	if err := database.DB.Select("*").Where("id = ?", user_id).Preload("Apps").First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// User data by id
func UserDataById(c *fiber.Ctx) error {

	var user models.User

	if err := database.DB.Select("id, name, surname, about, email, image").Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// Update name, surname and about
func UpdateUser(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return fiber.ErrInternalServerError
	}

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	if !utilities.OnlyEmptySpaces(data["name"]) {
		user.Name = data["name"]
	}

	if !utilities.OnlyEmptySpaces(data["mobile"]) {
		user.Name = data["mobile"]
	}

	if !utilities.OnlyEmptySpaces(data["surname"]) {
		user.Surname = data["surname"]
	}

	if !utilities.OnlyEmptySpaces(data["about"]) {
		user.About = data["about"]
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

// Delete user
func Delete(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Delete(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	user.DeleteAt = utilities.DateTimeNow()
	if err := database.DB.Save(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	utilities.UnauthorizedToken(token.Raw)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
