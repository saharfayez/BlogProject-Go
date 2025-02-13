package context

import (
	"goproject/database"
	postsimpl "goproject/impl/posts"
	usersimpl "goproject/impl/users"
	"goproject/interfaces/context"
	"goproject/interfaces/posts"
	"goproject/interfaces/users"
	"gorm.io/gorm"
)

type applicationContextImpl struct {
	propertiesConfig context.PropertiesConfig

	db *gorm.DB

	userRepository users.UserRepository

	userService users.UserService

	postRepository posts.PostRepository
}

func (Context *applicationContextImpl) GetPropertiesConfig() context.PropertiesConfig {
	return Context.propertiesConfig
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

	propertiesConf := newPropertiesConfig()

	appContext := &applicationContextImpl{propertiesConfig: propertiesConf}
	context.Context = appContext

	db, _ := database.InitDB()
	appContext.db = db

	userRepository := usersimpl.NewUserRepository(db)
	appContext.userRepository = userRepository

	userService := usersimpl.NewUserService(userRepository)
	appContext.userService = userService

	postRepository := postsimpl.NewPostRepository(db)
	appContext.postRepository = postRepository
}
