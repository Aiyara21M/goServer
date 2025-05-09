package employees

import (
	"server/models"
)

type EmployeeService interface {
	GetAllEmployees(dbName string) ([]models.EmployeeModel, error)
	GetEmployeeID(dbName string, id uint) (*models.EmployeeModel, error)
	CreateEmployee(dbName string, employee *models.EmployeeModel) error
}

type employeeService struct {
	repo EmployeeRepository
}

func (s *employeeService) GetAllEmployees(dbName string) ([]models.EmployeeModel, error) {
	return s.repo.FindAll(dbName)
}

func (s *employeeService) GetEmployeeID(dbName string, id uint) (*models.EmployeeModel, error) {
	return s.repo.FindByID(dbName, id)
}

func (s *employeeService) CreateEmployee(dbName string, employee *models.EmployeeModel) error {
	return s.repo.Create(dbName, employee)
}
