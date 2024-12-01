package routes

import (
	"github.com/labstack/echo/v4"
	"goproject/handlers"
)

func RegisterRoutes() *echo.Echo {
	router := echo.New()
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
	return router
}
