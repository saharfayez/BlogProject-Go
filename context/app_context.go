package context

import (
	"goproject/database"
	"goproject/posts"
	"goproject/users"
	"gorm.io/gorm"
	"log"
)

var Context ApplicationContext

type ApplicationContext interface {
	GetDB() *gorm.DB

	GetUserRepository() users.UserRepository

	GetUserService() users.UserService

	GetPostRepository() posts.PostRepository
}

type applicationContextImpl struct {
	db *gorm.DB

	userRepository users.UserRepository

	userService users.UserService

	postRepository posts.PostRepository
}

func (Context *applicationContextImpl) GetDB() *gorm.DB {
	return Context.db
}

func (Context *applicationContextImpl) GetUserRepository() users.UserRepository {
	return Context.userRepository
}

func (Context *applicationContextImpl) GetUserService() users.UserService {
	return Context.userService
}

func (Context *applicationContextImpl) GetPostRepository() posts.PostRepository {
	return Context.postRepository
}

func InitContext() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	postRepository := posts.NewPostRepository(db)

	Context = &applicationContextImpl{
		db,
		userRepository,
		userService,
		postRepository,
	}
}
