package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"
)

var JwtSecret = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
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

func GetRoleFromToken(claims *Claims) string {
	return claims.Role
}
