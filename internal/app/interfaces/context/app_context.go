package context

import (
	"goproject/internal/app/interfaces/posts"
	"goproject/internal/app/interfaces/users"
	"gorm.io/gorm"
)

var Context ApplicationContext

type ApplicationContext interface {
	GetPropertiesConfig() PropertiesConfig

	GetDB() *gorm.DB

	GetUserRepository() users.UserRepository

	GetUserService() users.UserService

	GetPostRepository() posts.PostRepository
}
