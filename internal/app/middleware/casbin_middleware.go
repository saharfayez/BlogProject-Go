package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"log"
)

func Enforce(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			role := GetRoleFromToken(c)
			action := c.Request().Method
			resource := c.Request().URL.Path

			fmt.Println("role", role, "action", action, "resource", resource)

			result, err := enforcer.Enforce(role, resource, action)
			if err != nil {
				log.Fatal("Error enforcing with username, resource, action: ", err)
			}

			fmt.Println("result:", result)

			if result {
				return next(c)
			}

			return echo.ErrForbidden
		}
	}
}
