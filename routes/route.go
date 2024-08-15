package routes

import (
	"github.com/fadiahaa/blog-go/controller"
	"github.com/fadiahaa/blog-go/middleware"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){
	// AUTH
	app.Post("/api/register",controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)

	// POST
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload-image", controller.Upload)
	app.Static("/api/uploads", "./uploads")
}