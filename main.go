package main

import (
	"github.com/labstack/echo/v4"
	"goproject/database"
	"net/http"
)

func main() {

	database.InitDB()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))

}
