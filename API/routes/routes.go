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

	userGroup.Use(middleware.RequireLogin)
	userGroup.Get("/idxpost", controllers.IndexPost)
	userGroup.Post("/createPost", controllers.CreatePost)
	userGroup.Get("/showpost", controllers.ShowPost)
	userGroup.Put("/updatepost/:id", controllers.UpdatePost)
	userGroup.Delete("/deletepost/:id", controllers.DeletePost)

	userGroup.Post("/like/:id", controllers.LikePost)
	userGroup.Delete("/deleteLike/:id", controllers.DeleteLikeDislike)
	userGroup.Post("/dislike/:id", controllers.DislikePost)

	userGroup.Get("/idxcomment", controllers.IndexComment)
	userGroup.Post("/comment/:id", controllers.CreateComment)
	userGroup.Put("/updatecomment/:id", controllers.UpdateComment)
	userGroup.Delete("/deletecomment/:id", controllers.DeleteComment)

	adminGroup := app.Group("/api/admin")
	adminGroup.Post("/create", controllers.CreateAdmin)
	adminGroup.Post("/loginadm", controllers.LoginAdmin)
	adminGroup.Get("/admin", controllers.GetAdmin)
	adminGroup.Post("/logout", controllers.LogoutAdm)

	adminGroup.Use(middleware.RequireLogin)
	adminGroup.Get("/edukasi", controllers.GetEdukasi)
	adminGroup.Post("/createedk", controllers.CreateEdukasi)
	adminGroup.Put("/updateedk/:id", controllers.UpdateEdukasi)
	adminGroup.Delete("/deleteedk/:id", controllers.DeleteEdukasi)

	adminGroup.Use(middleware.RequireLogin)
	adminGroup.Put("/updateuser/:id", controllers.UpdateStatusUser)
	adminGroup.Get("/getuser", controllers.GetUser)
}
