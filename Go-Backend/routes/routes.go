package routes

import (
	"Go-Backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/user", controllers.User)
	app.Post("/api/auth/logout", controllers.Logout)
	app.Post("/api/task/create", controllers.TaskCreate)
	app.Get("/api/task/list", controllers.TaskList)
}
