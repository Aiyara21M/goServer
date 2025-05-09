package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}

type userController struct {
	service UserService
}

func (ctrl *userController) GetUsers(c *fiber.Ctx) error {
	users, err := ctrl.service.GetUsers("db1")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := ctrl.service.GetUserByID("db1", uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}
