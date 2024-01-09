package routers

import (
	"auth_codex_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
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
	app.Post("/api/image_profile", controllers.ImageProfileUpload)
	app.Post("/api/image_wallpaper", controllers.ImageWallpaperUpload)
}
