package routers

import (
	"medium_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// create routers
	// static achives
	app.Static("/uploads", "./uploads")
	// auth routers
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Post("/api/auth/logout", controllers.Logout)
	// user
	app.Get("/api/user/profile", controllers.Profile)
	app.Get("/api/user/profile/:id", controllers.UserDataById)
	app.Delete("/api/user/profile", controllers.Delete)
	app.Post("/api/user/update", controllers.UpdateUser)
	// upload images
	app.Post("/api/imageProfile", controllers.ImageProfileUpload)
	app.Post("/api/imageWallpaper", controllers.ImageWallpaperUpload)
}