package main

import (
	"os"
	"server/apptcx"
	"server/config"
	"server/database"
	"server/models"
	"server/routes"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type UserModel struct {
	ID       int
	Username string
	Email    string
}

func main() {
	app := fiber.New()
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests – try again later.",
			})
		},
	}))

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	config.LoadEnvs(env)

	postgresConns, _ := database.ConnectAllPostgresDBs()
	mongoConns, _ := database.ConnectAllMongoDBs()

	apptcx.ConnectDB = &apptcx.AppContext{
		PostgresConnectors: postgresConns,
		MongoConnectors:    mongoConns,
	}

	postgresConns["db1"].GetDB().AutoMigrate(
		&models.UserModel{},
	)

	postgresConns["db3"].GetDB().AutoMigrate(
		&models.EmployeeModel{},
	)

	routes.SetupRoutes(app)

	app.Listen(":3011")
}
