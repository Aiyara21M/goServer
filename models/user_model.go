package models

type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:255;not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}

func (UserModel) TableName() string {
	return "Users"
}
