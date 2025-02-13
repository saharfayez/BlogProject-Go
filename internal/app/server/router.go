package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goproject/internal/app/business/impl/posts"
	"goproject/internal/app/business/impl/users"
	appMiddleware "goproject/internal/app/middleware"
)

func registerRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/signup", users.Signup)
	e.POST("/login", users.Login)

	e.Use(middleware.Logger())

	protected := e.Group("/api", appMiddleware.JWTMiddleware())

	protected.POST("/posts", posts.CreatePost)
	//protected.GET("/posts", posts.GetPosts)
	//protected.GET("/posts/:id", posts.GetPost)
	//protected.PUT("/posts/:id", posts.UpdatePost)
	//protected.DELETE("/posts/:id", posts.DeletePost)
	return e
}
