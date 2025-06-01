package users

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"goproject/internal/app/business/interfaces/users"
	middleware "goproject/internal/app/middleware"
	"goproject/internal/app/models"
	"log"
)

type userServiceImpl struct {
	userRepo users.UserRepository
}

func NewUserService(userRepo users.UserRepository) users.UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (userServiceImpl *userServiceImpl) FindUser(username string) (*models.User, error) {
	return userServiceImpl.userRepo.FindUserByUsername(username)
}

func (userServiceImpl *userServiceImpl) Signup(user *models.User) error {

	_, err := userServiceImpl.userRepo.FindUserByUsername(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}
	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashError != nil {
		return hashError
	}
	user.Password = string(hashedPassword)
	if saveErr := userServiceImpl.userRepo.Save(user); saveErr != nil {
		return saveErr
	}

	return nil
}

func (userServiceImpl *userServiceImpl) Login(username, password string) (string, error) {

	existingUser, err := userServiceImpl.userRepo.FindUserByUsername(username)
	if err != nil {
		log.Println("user service log:", err.Error())
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		log.Println("user service log passwords:", err.Error())
		return "", err
	}

	return middleware.GenerateJWT(username)
}

func MapUserDtoToUser(userDto UserDto) models.User {
	return models.User{
		Username: userDto.Username,
		Password: userDto.Password,
	}
}
