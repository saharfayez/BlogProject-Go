package mapping

import (
	"goproject/dtos"
	"goproject/models"
)

func MapUserDtoToUser(userDto dtos.UserDto) models.User {
	return models.User{
		Username: userDto.Username,
		Password: userDto.Password, // Password hashing should be handled separately
	}
}
