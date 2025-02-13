package context

import (
	"goproject/database"
	postsimpl "goproject/impl/posts"
	usersimpl "goproject/impl/users"
	"goproject/interfaces/context"
	"goproject/interfaces/posts"
	"goproject/interfaces/users"
	"gorm.io/gorm"
	"log"
)

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

// this is called once by go before main
func init() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := usersimpl.NewUserRepository(db)
	userService := usersimpl.NewUserService(userRepository)
	postRepository := postsimpl.NewPostRepository(db)

	context.Context = &applicationContextImpl{
		db,
		userRepository,
		userService,
		postRepository,
	}
}
