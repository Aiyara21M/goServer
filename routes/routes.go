package routes

import (
	"server/module/employees"
	"server/module/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	users.SetupUserRoutes(api.Group("/users"))

	employees.SetupEmployeeRoutes(api.Group("/employees"))
}
