package controllers

import (
	"fmt"
	"log"
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ImageProfileUpload(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)

	// parse incomming image file
	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./uploads dir
	err = c.SaveFile(file, fmt.Sprintf("./uploads/profile/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("%s", image)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	// remove image before save new image
	os.Remove("./uploads/profile/" + user.Image)

	// save url image in database
	user.Image = "/uploads/profile/" + imageUrl
	database.DB.Save(&user)
	// create meta data and send to client
	data := map[string]interface{}{
		"imageName": image,
		"imageUrl":  "/uploads/profile/" + imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func ImageWallpaperUpload(c *fiber.Ctx) error {
	// validate user
	token, err := utilities.IsAuthenticadToken(c, utilities.GoDotEnvVariable("SECRETKEY"))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)

	// parse incomming image file
	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./uploads dir
	err = c.SaveFile(file, fmt.Sprintf("./uploads/wallpaper/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("%s", image)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	// remove image before save new image
	os.Remove("./uploads/wallpaper/" + user.Image)

	// save url image in database
	user.Image = imageUrl
	database.DB.Save(&user)
	// create meta data and send to client
	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  "/uploads/wallpaper/" + imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
