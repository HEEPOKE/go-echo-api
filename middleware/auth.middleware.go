package middleware

import (
	"errors"
	"main/config"
	"main/services"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing JWT token")
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid JWT token")
			}
			return []byte(config.MainConfig().SECRET_KEY), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
		}

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Until(expirationTime) < 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "JWT token has expired")
		}

		userID, ok := claims["userID"].(float64)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
		}

		userService := services.NewUserService(&gorm.DB{})
		user, err := userService.GetUserByID(uint(userID))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
		}

		ctx.Set("user", user)
		return next(ctx)
	}
}
