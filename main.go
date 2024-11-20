package main

import (
	"goproject/database"
	"goproject/routes"
	"log"
	"net/http"
)

func main() {

	database.InitDB()

	router := routes.RegisterRoutes()

	log.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
