package services

import (
	"errors"
	"fmt"
	"log"
	"main/config"
	"main/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	secretKey []byte
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db, secretKey: []byte(config.MainConfig().SECRET_KEY)}
}

func (as *AuthService) RegisterUser(user *models.User) error {
	if as.db == nil {
		return fmt.Errorf("database is nil")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password for user %s: %s", user.Email, err.Error())
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	if err := as.db.Create(user).Error; err != nil {
		log.Printf("failed to create user %s: %s", user.Email, err.Error())
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (as *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := as.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (as *AuthService) GenerateToken(userID uint) (*models.Token, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(as.secretKey)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}, nil
}

func (as *AuthService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := as.db.Where(&models.User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}
