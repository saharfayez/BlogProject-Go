package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middlewares "goproject/middleware"
	"goproject/users"
	"goproject/posts"
)

func registerRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/signup", users.Signup)
	e.POST("/login", users.Login)

	e.Use(middleware.Logger())

	protected := e.Group("/api", middlewares.JWTMiddleware())

	protected.GET("/posts", posts.GetPosts)
	protected.POST("/posts", posts.CreatePost)
	protected.GET("/posts/:id", posts.GetPost)
	protected.PUT("/posts/:id", posts.UpdatePost)
	protected.DELETE("/posts/:id", posts.DeletePost)
	return e
}
