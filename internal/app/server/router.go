package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"goproject/internal/app/business/impl/posts"
	"goproject/internal/app/business/impl/users"
	appMiddleware "goproject/internal/app/middleware"
	"log"
)

func registerRoutes() *echo.Echo {

	e := echo.New()

	e.POST("/signup", users.Signup)
	e.POST("/login", users.Login)

	// zap
	//zapLogger, _ := zap.NewProduction()
	//e.Use(appMiddleware.ZapLoggerMiddleware(zapLogger))

	// slog
	//slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//e.Use(appMiddleware.SlogLoggerMiddleware(slogLogger))

	//e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//	c.Set("responseBody", string(resBody))
	//}))

	protected := e.Group("/api")

	protected.Use(appMiddleware.JWTMiddleware())

	enforcer, err := casbin.NewEnforcer("../../internal/app/server/model.conf", "../../internal/app/server/policy.csv")
	if err != nil {
		log.Fatal("Failed to create Casbin enforcer: ", err)
	}

	protected.Use(appMiddleware.Enforce(enforcer))

	protected.POST("/posts", posts.CreatePost)

	//protected.GET("/posts", posts.GetPosts)
	//protected.GET("/posts/:id", posts.GetPost)
	//protected.PUT("/posts/:id", posts.UpdatePost)
	//protected.DELETE("/posts/:id", posts.DeletePost)
	return e
}
