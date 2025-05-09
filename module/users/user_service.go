package users

import "server/models"

type UserService interface {
	GetUsers(dbName string) ([]models.UserModel, error)
	GetUserByID(dbName string, id uint) (*models.UserModel, error)
}

type userService struct {
	repo UserRepository
}

func (s *userService) GetUsers(dbName string) ([]models.UserModel, error) {

	return s.repo.FindAll(dbName)
}

func (s *userService) GetUserByID(dbName string, id uint) (*models.UserModel, error) {
	return s.repo.FindByID(dbName, id)
}
