package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goproject/handlers"
	middlewares "goproject/middleware"
)

func RegisterRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)

	e.Use(middleware.Logger())

	protected := e.Group("/api", middlewares.JWTMiddleware())

	protected.GET("/posts", handlers.GetPosts)
	protected.POST("/posts", handlers.CreatePost)
	protected.GET("/posts/:id", handlers.GetPost)
	protected.PUT("/posts/:id", handlers.UpdatePost)
	protected.DELETE("/posts/:id", handlers.DeletePost)
	return e
}
