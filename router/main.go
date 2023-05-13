package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vladyslav-dev/todo-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)

	todos := app.Group("/todos")
	todos.Get("/", handlers.GetAllTodos)
	todos.Post("/", handlers.CreateTodo)
	todos.Put("/:id", handlers.UpdateTodo)
	todos.Get("/:id", handlers.GetOneTodo)
	todos.Delete("/:id", handlers.DeleteTodo)
}
