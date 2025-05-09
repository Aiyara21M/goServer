package employees

import (
	"server/apptcx"
	"server/models"
)

type EmployeeRepository interface {
	FindAll(dbName string) ([]models.EmployeeModel, error)
	FindByID(dbName string, id uint) (*models.EmployeeModel, error)
	Create(dbName string, employee *models.EmployeeModel) error
}

type employeeRepository struct{}

func (r *employeeRepository) FindAll(dbName string) ([]models.EmployeeModel, error) {
	db := apptcx.ConnectDB.PostgresConnectors[dbName].GetDB()
	var employees []models.EmployeeModel
	err := db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) FindByID(dbName string, id uint) (*models.EmployeeModel, error) {
	db := apptcx.ConnectDB.PostgresConnectors[dbName].GetDB()
	var employee models.EmployeeModel
	err := db.First(&employee, id).Error
	return &employee, err
}

func (r *employeeRepository) Create(dbName string, employee *models.EmployeeModel) error {
	db := apptcx.ConnectDB.PostgresConnectors[dbName].GetDB()
	return db.Create(employee).Error
}
