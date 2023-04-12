package routes

import (
	"main/config"
	"main/controllers"
	jwtMiddleware "main/middleware"
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
	api.POST("/auth/logout", authController.Logout, jwtMiddleware.JWTMiddleware)

	userController := &controllers.UserController{}
	api.GET("/users", userController.GetAllUsers, jwtMiddleware.JWTMiddleware)
	api.POST("/users/create", userController.CreateUser, jwtMiddleware.JWTMiddleware)
	api.GET("/users/:id", userController.GetUserByID, jwtMiddleware.JWTMiddleware)
	api.PUT("/users/update/:id", userController.UpdateUser, jwtMiddleware.JWTMiddleware)
	api.DELETE("/users/delete/:id", userController.DeleteUser, jwtMiddleware.JWTMiddleware)

	e.Logger.Fatal(e.Start(":" + config.MainConfig().PORT))
}
