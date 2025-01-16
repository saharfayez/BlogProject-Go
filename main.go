package main

import (
	"fmt"
	"goproject/database"
	"goproject/server"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())
	database.InitDB()
	server.Serve()
}
