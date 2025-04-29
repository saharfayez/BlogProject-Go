package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RBACMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			requestMethod := c.Request().Method
			requestPath := c.Path()

			token := GetTokenFromContext(c)
			roles := GetRolesFromToken(token)

			for _, role := range roles {
				for _, permission := range role.Permissions {
					for _, resource := range permission.Resources {
						if permission.Name == requestMethod && resource.Name == requestPath {
							return next(c)
						}
					}
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden: You don't have access")
		}
	}
}
