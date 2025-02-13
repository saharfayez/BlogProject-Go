package users

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"goproject/internal/app/context"
	middleware "goproject/internal/app/middleware"
	"goproject/internal/app/models"
	"net/http"
)

func Signup(c echo.Context) error {

	var userDto UserDto
	if err := c.Bind(&userDto); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := MapUserDtoToUser(userDto)

	userService := context.Context.GetUserService()

	err := userService.Signup(&user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var signupResponse SignUpResponseDto

	signupResponse.ID = user.ID
	signupResponse.Username = user.Username

	return c.JSON(http.StatusCreated, signupResponse)
}

// TODO - refactor to use UserService
func Login(c echo.Context) error {

	var userDto UserDto
	if err := c.Bind(&userDto); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var user models.User
	result := context.Context.GetDB().Where("username = ?", userDto.Username).First(&user)
	if result.Error != nil {
		c.Logger().Error(result.Error)
		return c.String(http.StatusNotFound, "User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusUnauthorized, "Passwords are not compatible")
	}

	token, err := middleware.GenerateJWT(user.Username)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Error generating token")
	}

	var loginResponse LoginResponseDto
	loginResponse.Token = token

	return c.JSON(http.StatusOK, loginResponse)
}
