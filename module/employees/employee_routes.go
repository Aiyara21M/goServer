package employees

import (
	"server/secure"

	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(router fiber.Router) {

	repo := &employeeRepository{}
	service := &employeeService{repo: repo}
	ctrl := &employeeController{service: service}

	router.Get("/", ctrl.GetEmployees)
	router.Post("/", secure.VerifyHMACMiddleware, ctrl.GetEmployeeByID)
	// router.Post("/create", ctrl.CreateSampleEmployee)
}
