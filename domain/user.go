package domain

import (
	"time"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, role, email string, password string) *User {
	return &User{
		Name:      name,
		Role:      role,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}

// TODO
func EncryptPassword() {}
