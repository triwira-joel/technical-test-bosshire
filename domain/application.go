package domain

import "time"

type CreateApplicationRequest struct {
	JobID      int `json:"job_id"`
	TalentID   int `json:"talent_id"`
	EmployerID int `json:"employer_id"`
}

type UpdateApplicationRequest struct {
	Status string `json:"status"`
}

type Application struct {
	ID         int       `json:"id"`
	JobID      int       `json:"job_id"`
	TalentID   int       `json:"talent_id"`
	EmployerID int       `json:"employer_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewApplication(job_id, talent_id, employer_id int, status string) *Application {
	return &Application{
		JobID:      job_id,
		TalentID:   talent_id,
		EmployerID: employer_id,
		Status:     status,
		CreatedAt:  time.Now().UTC(),
	}
}

type status string

var (
	Review    status = "REVIEW"
	Interview status = "INTERVIEW"
	Reject    status = "REJECT"
	Accept    status = "ACCEPT"
)
