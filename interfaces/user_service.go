package interfaces

import "goproject/models"

type UserService interface {
	Signup(user *models.User) error
	Login(username, password string) (string, error)
}
