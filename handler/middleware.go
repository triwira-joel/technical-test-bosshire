package handler

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func CreateJWTToken(name, email, role string) (string, error) {
	claims := JWTClaims{
		name,
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("MySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
