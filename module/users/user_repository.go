package users

import (
	"server/apptcx"
	"server/models"
)

type UserRepository interface {
	FindAll(dbName string) ([]models.UserModel, error)
	FindByID(dbName string, id uint) (*models.UserModel, error)
}

type userRepository struct{}

func (r *userRepository) FindAll(dbName string) ([]models.UserModel, error) {
	users := []models.UserModel{}

	db := apptcx.ConnectDB.PostgresConnectors[dbName].GetDB()
	err := db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(dbName string, id uint) (*models.UserModel, error) {
	user := models.UserModel{}
	db := apptcx.ConnectDB.PostgresConnectors[dbName].GetDB()
	err := db.First(&user, id).Error
	return &user, err
}
