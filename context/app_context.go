package context

import (
	"goproject/database"
	"goproject/posts"
	"goproject/users"
	"gorm.io/gorm"
	"log"
)

type ApplicationContext struct {
	DB *gorm.DB

	UserRepository users.UserRepository

	UserService    users.UserService

	PostRepository posts.PostRepository
}

var Context *ApplicationContext

func InitContext() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	postRepository := posts.NewPostRepository(db)

	Context = &ApplicationContext{
		db,
		userRepository,
		userService,
		postRepository,
	}
}
