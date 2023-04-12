package services

import (
	"errors"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return as.db.Create(user).Error
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
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
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
