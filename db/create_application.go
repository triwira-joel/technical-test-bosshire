package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) CreateApplication(a *domain.Application) error {
	query := `
		INSERT INTO applications (job_id, talent_id, employer_id, status, created_at) 
			VALUES ($1, $2, $3, $4, $5) 
				RETURNING id;`

	_, err := s.DB.Exec(query, a.JobID, a.TalentID, a.EmployerID, a.Status, a.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
