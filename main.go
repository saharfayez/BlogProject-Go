package main

import (
	"fmt"
	"goproject/database"
	"goproject/route"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())
	database.InitDB()
	e := routes.RegisterRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
