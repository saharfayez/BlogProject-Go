package context

import (
	"goproject/interfaces/posts"
	"goproject/interfaces/users"
	"gorm.io/gorm"
)

var Context ApplicationContext

type ApplicationContext interface {
	GetDB() *gorm.DB

	GetUserRepository() users.UserRepository

	GetUserService() users.UserService

	GetPostRepository() posts.PostRepository
}
