package routes

import (
	"main/config"
	"main/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("/api")

	authController := &controllers.AuthController{}
	api.POST("/auth/login", authController.Login)
	api.POST("/auth/register", authController.Register)
	api.POST("/auth/logout", authController.Logout)

	userController := &controllers.UserController{}
	api.GET("/users", userController.GetAllUsers)
	api.POST("/users/create", userController.CreateUser)
	api.GET("/users/:id", userController.GetUserByID)
	api.PUT("/users/update/:id", userController.UpdateUser)
	api.DELETE("/users/delete/:id", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":" + config.MainConfig().PORT))
}
