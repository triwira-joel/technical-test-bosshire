package db

import "github.com/triwira-joel/technical-test-bosshire/domain"

func (s *DBHandler) SelectApplicationsByEmployerID(id int) ([]*domain.Application, error) {
	rows, err := s.DB.Query("SELECT * FROM applications WHERE employer_id = $1", id)
	if err != nil {
		return nil, err
	}

	applications := []*domain.Application{}
	for rows.Next() {
		application := new(domain.Application)
		err := rows.Scan(
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

		applications = append(applications, application)
	}

	return applications, nil
}
