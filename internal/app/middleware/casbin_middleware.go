package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Enforce(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := GetTokenFromContext(c)
			role := GetRoleFromToken(token)
			action := c.Request().Method
			resource := c.Request().URL.Path

			fmt.Println("role", role, "action", action, "resource", resource)

			result, err := enforcer.Enforce(role, resource, action)
			if err != nil {
				log.Fatal("Error enforcing with username, resource, action: ", err)
			}

			if result {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusForbidden, "Forbidden: You don't have access")

		}
	}
}
