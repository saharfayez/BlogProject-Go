package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: JwtSecret,
		ContextKey: "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &Claims{}
		},
	})
}
