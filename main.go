package main

import (
	"goproject/database"
	"goproject/route"
)

func main() {

	database.InitDB()
	e := routes.RegisterRoutes()
	e.Logger.Fatal(e.Start(":8080"))

}
