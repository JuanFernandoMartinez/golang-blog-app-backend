package server

import (
	"log"

	"example.com/blog-app/controllers"
	"example.com/blog-app/initializers"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	//user routes
	app.Post("/user", controllers.CreateUser)
	app.Get("/user", controllers.GetAllUsers)
	app.Get("/user/:id", controllers.GetUserByID)
	app.Put("/user/:id", controllers.EditUser)
	app.Delete("/user/:id", controllers.DeleteUser)

	//thought routes
	app.Post("/thought", controllers.CreateThought)
	app.Get("/thought", controllers.GetAllThoughts)
	app.Get("/thought/:id", controllers.GetThoughtByID)
	app.Put("/thought/:id", controllers.EditThought)
	app.Delete("/thought/:id", controllers.DeleteThought)

	//comment routes
	app.Post("/comment", controllers.CreateComment)
	app.Get("/comment", controllers.GetAllComments)
	app.Get("/comment/:id", controllers.GetCommentByID)
	app.Put("/comment/:id", controllers.EditComment)
	app.Delete("/comment/:id", controllers.DeleteComment)

	log.Fatal(app.Listen(":" + initializers.SERVER_PORT))
}
