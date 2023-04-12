package controllers

import (
	"errors"
	"main/models"
	"main/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Register(ctx echo.Context) error {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}
	err = ac.authService.RegisterUser(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create user",
		})
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (ac *AuthController) Login(ctx echo.Context) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&loginData); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}

	user, err := ac.authService.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"error": "user not found",
			})
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid password",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to authenticate user",
		})
	}

	token, err := ac.authService.GenerateToken(user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to generate JWT token",
		})
	}

	return ctx.JSON(http.StatusOK, token)
}

func (ac *AuthController) Logout(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
