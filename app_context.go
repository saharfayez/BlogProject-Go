package main

import (
	"goproject/database"
	"goproject/interfaces"
	"goproject/posts"
	"goproject/users"
	"gorm.io/gorm"
	"log"
)

type applicationContextImpl struct {
	db *gorm.DB

	userRepository interfaces.UserRepository

	userService interfaces.UserService

	postRepository interfaces.PostRepository
}

func (Context *applicationContextImpl) GetDB() *gorm.DB {
	return Context.db
}

func (Context *applicationContextImpl) GetUserRepository() interfaces.UserRepository {
	return Context.userRepository
}

func (Context *applicationContextImpl) GetUserService() interfaces.UserService {
	return Context.userService
}

func (Context *applicationContextImpl) GetPostRepository() interfaces.PostRepository {
	return Context.postRepository
}

// this is called once by go before main
func init() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	postRepository := posts.NewPostRepository(db)

	interfaces.Context = &applicationContextImpl{
		db,
		userRepository,
		userService,
		postRepository,
	}
}
