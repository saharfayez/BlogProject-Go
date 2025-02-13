package users

import "goproject/internal/app/models"

type UserRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	Save(user *models.User) error
}
