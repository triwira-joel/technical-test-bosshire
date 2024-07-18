package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) CreateUser(u *domain.User) error {
	query := `INSERT INTO users (name, role, email, password, created_at) VALUES ($1, $2, $3, $4, $5);`

	if _, err := s.DB.Exec(query, u.Name, u.Role, u.Email, u.Password, u.CreatedAt); err != nil {
		return err
	}

	return nil
}
