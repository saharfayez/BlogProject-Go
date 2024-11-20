package routes

import (
	"github.com/gorilla/mux"
	"goproject/handlers"
	"goproject/middlewares"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middlewares.JWTMiddleware)
	protected.HandleFunc("/posts", handlers.GetPosts()).Methods("GET")
	protected.HandleFunc("/posts", handlers.CreatePost()).Methods("POST")
	protected.HandleFunc("/posts/{id}", handlers.GetPost()).Methods("GET")
	protected.HandleFunc("/posts/{id}", handlers.UpdatePost()).Methods("PUT")
	protected.HandleFunc("/posts/{id}", handlers.DeletePost()).Methods("DELETE")

	return router
}
