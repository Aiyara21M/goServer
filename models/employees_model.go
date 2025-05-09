package models

type EmployeeModel struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:255;not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}

func (EmployeeModel) TableName() string {
	return "Employees"
}
