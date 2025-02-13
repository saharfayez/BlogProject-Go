package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"goproject/interfaces"
	"goproject/models"
	"goproject/utils"
)

type userServiceImpl struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (userServiceImpl *userServiceImpl) GetName() string {
	return "UserService"
}

func (userServiceImpl *userServiceImpl) Signup(user *models.User) error {

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

func (userServiceImpl *userServiceImpl) Login(username, password string) (string, error) {

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

func MapUserDtoToUser(userDto UserDto) models.User {
	return models.User{
		Username: userDto.Username,
		Password: userDto.Password, // Password hashing should be handled separately
	}
}
