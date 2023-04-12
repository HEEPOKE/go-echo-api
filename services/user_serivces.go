package services

import (
	"errors"
	"main/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where(&models.User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	result := s.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) DeleteUser(user *models.User) error {
	result := s.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
