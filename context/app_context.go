package context

import (
	"goproject/database"
	"goproject/posts"
	"goproject/users"
	"gorm.io/gorm"
	"log"
)

type ApplicationContext interface {
	GetDB() *gorm.DB

	GetUserRepository() users.UserRepository

	GetUserService() users.UserService

	GetPostRepository() posts.PostRepository
}

type ApplicationContextImpl struct {
	db *gorm.DB

	userRepository users.UserRepository

	userService users.UserService

	postRepository posts.PostRepository
}

func (Context *ApplicationContextImpl) GetDB() *gorm.DB {
	return Context.db
}

func (Context *ApplicationContextImpl) GetUserRepository() users.UserRepository {
	return Context.userRepository
}

func (Context *ApplicationContextImpl) GetUserService() users.UserService {
	return Context.userService
}

func (Context *ApplicationContextImpl) GetPostRepository() posts.PostRepository {
	return Context.postRepository
}

var Context ApplicationContext

func InitContext() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	postRepository := posts.NewPostRepository(db)

	Context = &ApplicationContextImpl{
		db,
		userRepository,
		userService,
		postRepository,
	}
}
