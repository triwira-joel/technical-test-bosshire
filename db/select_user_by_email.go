package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) SelectUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := s.DB.QueryRow(`SELECT * FROM users WHERE email = $1;`, email).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
