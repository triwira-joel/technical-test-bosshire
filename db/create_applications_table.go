package db

func (s *DBHandler) CreateApplicationsTable() error {
	query := `CREATE TABLE IF NOT EXISTS applications (
		id SERIAL PRIMARY KEY,
		job_id INT NOT NULL,
		talent_id INT NOT NULL,
		employer_id INT NOT NULL,
		status VARCHAR(128) NOT NULL,
		created_at TIMESTAMP NOT NULL,

		CONSTRAINT fk_job FOREIGN KEY(job_id) REFERENCES jobs(id),
		CONSTRAINT fk_talent FOREIGN KEY(talent_id) REFERENCES users(id),
		CONSTRAINT fk_employer FOREIGN KEY(employer_id) REFERENCES users(id)
	);`

	_, err := s.DB.Exec(query)
	return err
}
