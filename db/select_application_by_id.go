package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) SelectApplicationByID(id int) (*domain.Application, error) {
	var application domain.Application

	err := s.DB.QueryRow(`SELECT * FROM applications WHERE id = $1;`, id).Scan(
		&application.ID,
		&application.JobID,
		&application.TalentID,
		&application.EmployerID,
		&application.Status,
		&application.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &application, nil
}
