package db

import (
	"github.com/triwira-joel/technical-test-bosshire/domain"
)

func (s *DBHandler) UpdateApplicationStatus(a *domain.Application) error {
	query := `UPDATE applications SET status = $1 WHERE id = $2;`

	_, err := s.DB.Exec(query, a.Status, a.ID)
	if err != nil {
		return err
	}

	return nil
}
