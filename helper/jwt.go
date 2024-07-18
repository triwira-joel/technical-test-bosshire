package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

type JWTClaims struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func CreateJWTToken(user *domain.User) (string, error) {
	claims := JWTClaims{
		user.Name,
		user.Id,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour).Unix(),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rt.SignedString([]byte("MySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func IsAuthorized(requestToken string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("MySecret"), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractDataFromToken(requestToken string) (int, string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("MySecret"), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	return int(claims["id"].(float64)), claims["role"].(string), nil
}
