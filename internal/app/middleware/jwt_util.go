package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"goproject/internal/app/models"
	"time"
)

var JwtSecret = []byte("your_secret_key")

type Claims struct {
	Username string        `json:"username"`
	Roles    []models.Role `json:"roles"`
	jwt.RegisteredClaims
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}

func GenerateJWT(username string, roles []models.Role) (string, error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &Claims{
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func GetTokenFromContext(c echo.Context) *Claims {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*Claims)
	return claims
}

func GetUsernameFromToken(claims *Claims) string {
	return claims.Username
}

func GetRolesFromToken(claims *Claims) []models.Role {
	return claims.Roles
}
