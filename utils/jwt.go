package utils

import (
	"fmt"
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

func GetTokenFromContext(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*Claims)
	return claims.Username
}

func GetUsernameFromToken(tokenString string) (string, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.Username, nil
}
