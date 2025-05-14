package services

import (
	"go_backend/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return models.GetAllUser(us.db)
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return models.GetUserByID(us.db, id)
}
