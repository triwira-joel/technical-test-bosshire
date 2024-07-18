package db

func (s *DBHandler) CreateUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}
