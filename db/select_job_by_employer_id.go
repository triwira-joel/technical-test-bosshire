package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) SelectJobByEmployerID(id int) ([]*domain.Job, error) {
	rows, err := s.DB.Query("SELECT * FROM jobs WHERE employer_id = $1", id)
	if err != nil {
		return nil, err
	}

	jobs := []*domain.Job{}
	for rows.Next() {
		job := new(domain.Job)
		err := rows.Scan(
			&job.Id,
			&job.Name,
			&job.Description,
			&job.EmployerID,
			&job.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}
