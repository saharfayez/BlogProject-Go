package users

import (
	"github.com/labstack/echo/v4"
	"goproject/internal/app/context"
	"net/http"
)

func Signup(c echo.Context) error {

	var userDto UserDto
	if err := c.Bind(&userDto); err != nil {
		_ = c.String(http.StatusBadRequest, err.Error())
		return err
	}
	user := MapUserDtoToUser(userDto)

	userService := context.Context.GetUserService()

	err := userService.Signup(&user)
	if err != nil {
		_ = c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	var signupResponse SignUpResponseDto

	signupResponse.ID = user.ID
	signupResponse.Username = user.Username

	return c.JSON(http.StatusCreated, signupResponse)
}

func Login(c echo.Context) error {

	var userDto UserDto
	if err := c.Bind(&userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userService := context.Context.GetUserService()

	token, err := userService.Login(userDto.Username, userDto.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	loginResponse := LoginResponseDto{
		Token: token,
	}

	return c.JSON(http.StatusOK, loginResponse)
}
