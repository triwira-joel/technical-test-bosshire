package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) CreateJob(j *domain.Job) error {
	query := `
		INSERT INTO jobs (name, description, employer_id, created_at) 
			VALUES ($1, $2, $3, $4) 
				RETURNING id;`

	_, err := s.DB.Exec(query, j.Name, j.Description, j.EmployerID, j.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
