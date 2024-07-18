package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) SelectUserByID(id int) (*domain.User, error) {
	var user domain.User

	err := s.DB.QueryRow(`SELECT * FROM users WHERE id = $1;`, id).Scan(
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
