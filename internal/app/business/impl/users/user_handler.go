package users

import (
	"github.com/labstack/echo/v4"
	"goproject/internal/app/context"
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

func Login(c echo.Context) error {

	var userDto UserDto
	if err := c.Bind(&userDto); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userService := context.Context.GetUserService()

	token, err := userService.Login(userDto.Username, userDto.Password)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Error generating token")
	}

	loginResponse := LoginResponseDto{
		Token: token,
	}

	return c.JSON(http.StatusOK, loginResponse)
}
