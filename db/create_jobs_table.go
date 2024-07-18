package db

func (s *DBHandler) CreateJobsTable() error {
	query := `CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		employer_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL,

		CONSTRAINT fk_user FOREIGN KEY(employer_id) REFERENCES users(id)
	);`

	_, err := s.DB.Exec(query)
	return err
}
