package interfaces

import (
	"gorm.io/gorm"
)

var Context ApplicationContext

type ApplicationContext interface {
	GetDB() *gorm.DB

	GetUserRepository() UserRepository

	GetUserService() UserService

	GetPostRepository() PostRepository
}
