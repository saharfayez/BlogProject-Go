package users

import "goproject/internal/app/models"

type UserService interface {
	Signup(user *models.User) error
	Login(username, password string) (string, error)
	FindUser(username string) (*models.User, error)
}
