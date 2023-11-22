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

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	if err := database.DB.Select("id, name, surname, about, date_member, email, mobile, image_profile").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// User data by id
func UserDataById(c *fiber.Ctx) error {

	// _, err := utilities.IsAuthenticadToken(c, SecretKey)

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	var user models.User

	if err := database.DB.Select("id, name, surname, about, date_member, email, image_profile").Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// Update name, surname and about
func UpdateUser(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

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

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

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

	utilities.UnauthorizedToken(token.Raw)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
