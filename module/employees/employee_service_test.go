// employee_service_test.go
package employees // <<< เปลี่ยนตรงนี้

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEmployeeRepo struct{}

func (m *mockEmployeeRepo) FindAll(dbName string) ([]models.EmployeeModel, error) { return nil, nil }
func (m *mockEmployeeRepo) FindByID(dbName string, id uint) (*models.EmployeeModel, error) {
	return &models.EmployeeModel{ID: id, Username: "John", Email: "john@example.com"}, nil
}
func (m *mockEmployeeRepo) Create(dbName string, employee *models.EmployeeModel) error { return nil }

func TestGetEmployeeID(t *testing.T) {
	repo := &mockEmployeeRepo{}
	service := &employeeService{repo: repo}

	employee, err := service.GetEmployeeID("db3", 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), employee.ID)
	assert.Equal(t, "John", employee.Username)
}
