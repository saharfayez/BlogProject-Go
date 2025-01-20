package main

import (
	"fmt"
	"goproject/database"
	"goproject/server"
	"log"
	"runtime"
)

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version())
	_, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	server.Serve()
}
