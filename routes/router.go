package routes

import (
	"main/config"
	"main/controllers"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func Router() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	configJwt := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("my-secret-key"),
	}
	jwtMiddleware := middleware.JWTWithConfig(configJwt)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HUI IS SMART!")
	})

	api := e.Group("/api")

	authController := &controllers.AuthController{}
	api.POST("/auth/login", authController.Login)
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/logout", authController.Logout, jwtMiddleware)

	userController := &controllers.UserController{}
	api.GET("/users", userController.GetAllUsers, jwtMiddleware)
	api.POST("/users/create", userController.CreateUser, jwtMiddleware)
	api.GET("/users/:id", userController.GetUserByID, jwtMiddleware)
	api.PUT("/users/update/:id", userController.UpdateUser, jwtMiddleware)
	api.DELETE("/users/delete/:id", userController.DeleteUser, jwtMiddleware)

	e.Logger.Fatal(e.Start(":" + config.MainConfig().PORT))
}
