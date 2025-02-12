package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"goproject/models"
	"goproject/repository"
	"goproject/utils"
)

type UserService interface {
	Service
	Signup(user *models.User) error
	Login(username, password string) (string, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (userServiceImpl *UserServiceImpl) GetName() string {
	return "UserService"
}

func (userServiceImpl *UserServiceImpl) Signup(user *models.User) error {

	existingUser, _ := userServiceImpl.userRepo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("User already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return userServiceImpl.userRepo.Save(user)
}

func (userServiceImpl *UserServiceImpl) Login(username, password string) (string, error) {

	existingUser, err := userServiceImpl.userRepo.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(existingUser.Password))
	if err != nil {
		return "", err
	}

	return utils.GenerateJWT(username)
}
