package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"
)

var JwtSecret = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func GetUsernameFromContext(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*Claims)
	username := claims.Username
	return username
}
