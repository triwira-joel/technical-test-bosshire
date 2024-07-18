package db

import "github.com/triwira-joel/technical-test-bosshire/domain"

func (s *DBHandler) SelectUsers() ([]*domain.User, error) {
	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := []*domain.User{}
	for rows.Next() {
		user := new(domain.User)
		err := rows.Scan(
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

		users = append(users, user)
	}

	return users, nil
}
