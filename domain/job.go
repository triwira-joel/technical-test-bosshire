package domain

import "time"

type CreateJobRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EmployerID  int    `json:"employer_id"`
}

type Job struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EmployerID  int       `json:"employer_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewJob(name string, description string, employer_id int) *Job {
	return &Job{
		Name:        name,
		Description: description,
		EmployerID:  employer_id,
		CreatedAt:   time.Now().UTC(),
	}
}
