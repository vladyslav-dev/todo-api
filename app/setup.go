package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vladyslav-dev/todo-api/config"
	"github.com/vladyslav-dev/todo-api/database"
	"github.com/vladyslav-dev/todo-api/router"
)

func SetupAndRunApp() error {
	// Load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// Start database
	err = database.ConnectMongoDB()
	if err != nil {
		return err
	}

	// Defer closing database
	defer database.DisconnectMongoDB()

	// Create App
	app := fiber.New()

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// Routes
	router.SetupRoutes(app)

	// Swagger

	// Run App
	app.Listen(":" + os.Getenv("PORT"))

	return nil
}
