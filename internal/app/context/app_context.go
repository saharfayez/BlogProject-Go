package context

import (
	"goproject/internal/app/business/interfaces/posts"
	"goproject/internal/app/business/interfaces/users"
	"gorm.io/gorm"
)

var Context ApplicationContext

type ApplicationContext interface {
	GetPropertiesConfig() PropertiesConfig

	GetDB() *gorm.DB

	GetUserRepository() users.UserRepository

	GetUserService() users.UserService

	GetPostRepository() posts.PostRepository

	GetPostService() posts.PostService
}
