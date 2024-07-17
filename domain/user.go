package domain

import "time"

type CreateUserRequest struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, role, email string) *User {
	return &User{
		Name:      name,
		Role:      role,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
}
