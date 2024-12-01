package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"goproject/database"
	"goproject/dtos"
	"goproject/mapping"
	"goproject/models"
	"goproject/response"
	"goproject/utils"
	"net/http"
)

func Signup(c echo.Context) error {

	var userDto dtos.UserDto
	if err := c.Bind(&userDto); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := mapping.MapUserDtoToUser(userDto)
	var existingUser models.User

	error := database.DB.Where("username = ?", user.Username).First(&existingUser).Error

	if error == nil {
		c.Logger().Error(error)
		return c.String(http.StatusBadRequest, "Username already exists. Please choose another one.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Error hashing password")

	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.Logger().Error(result.Error)
		return c.String(http.StatusInternalServerError, "Error creating user")

	}
	var signupResponse response.SignUpResponse

	signupResponse.ID = user.ID
	signupResponse.Username = user.Username

	return c.JSON(http.StatusCreated, signupResponse)
}

func Login(c echo.Context) error {

	var userDto dtos.UserDto
	if err := c.Bind(&userDto); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var user models.User
	result := database.DB.Where("username = ?", userDto.Username).First(&user)
	if result.Error != nil {
		c.Logger().Error(result.Error)
		return c.String(http.StatusNotFound, "User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusUnauthorized, "Passwords are not compatible")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Error generating token")
	}
	var loginResponse response.LoginResponse
	loginResponse.Token = token
	return c.JSON(http.StatusOK, loginResponse)
}
