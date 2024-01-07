package routes

import (
	"petani_edukasi/controllers"
	"petani_edukasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	userGroup := app.Group("/api/user")
	userGroup.Post("/register", controllers.Register)
	userGroup.Post("/login", controllers.Login)
	userGroup.Get("/profile", controllers.User)
	userGroup.Post("/logout", controllers.Logout)

	adminGroup := app.Group("/api/admin")
	adminGroup.Post("/create", controllers.CreateAdmin)
	adminGroup.Post("/loginadm", controllers.LoginAdmin)
	adminGroup.Get("/admin", controllers.Admin)
	adminGroup.Post("/logout", controllers.LogoutAdm)

	adminGroup.Use(middleware.RequireLogin)
	adminGroup.Get("/edukasi", controllers.GetEdukasi)
	adminGroup.Post("/createedk", controllers.CreateEdukasi)
	adminGroup.Put("/updateedk/:id", controllers.UpdateEdukasi)
	adminGroup.Delete("/deleteedk/:id", controllers.DeleteEdukasi)
}
