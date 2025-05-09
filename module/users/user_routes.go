package users

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {

	repo := &userRepository{}
	service := &userService{repo: repo}
	ctrl := &userController{service: service}

	router.Get("/", ctrl.GetUsers)
	router.Get("/:id", ctrl.GetUserByID)
}
