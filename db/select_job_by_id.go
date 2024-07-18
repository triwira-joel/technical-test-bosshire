package db

import (
	"fmt"

	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) SelectJobByID(id int) (*domain.Job, error) {
	var job domain.Job

	fmt.Println("-- DB -- ", id)

	err := s.DB.QueryRow(`SELECT * FROM jobs WHERE id = $1;`, id).Scan(
		&job.Id,
		&job.Name,
		&job.Description,
		&job.EmployerID,
		&job.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
