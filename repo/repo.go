package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/triwira-joel/technical-test-bosshire/db"
	d "github.com/triwira-joel/technical-test-bosshire/domain"
)

type Repo struct {
	db *db.DBHandler
}

func NewRepo(
	simpledb *db.DBHandler,
) *Repo {
	return &Repo{
		simpledb,
	}
}

func (r *Repo) GetUser(c echo.Context, id int) (*d.User, error) {
	user, err := r.db.SelectUserByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) CreateUser(c echo.Context, name string, role string, email string, password string) (*d.User, error) {
	newUser := d.NewUser(name, role, email, password)

	if err := r.db.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *Repo) GetUserByEmailAndPassword(c echo.Context, email, password string) (*d.User, error) {
	user, err := r.db.SelectUserByEmailAndPassword(email, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetUserByEmail(c echo.Context, email string) (*d.User, error) {
	user, err := r.db.SelectUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repo) GetUsers(c echo.Context) ([]*d.User, error) {
	users, err := r.db.SelectUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repo) GetJobs(c echo.Context) ([]*d.Job, error) {
	jobs, err := r.db.SelectJobs()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *Repo) CreateJob(c echo.Context, name string, description string, employer_id int) (*d.Job, error) {
	newJob := d.NewJob(name, description, employer_id)

	if err := r.db.CreateJob(newJob); err != nil {
		return nil, err
	}

	return newJob, nil
}

func (r *Repo) GetJob(c echo.Context, id int) (*d.Job, error) {
	job, err := r.db.SelectJobByID(id)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (r *Repo) GetJobsByEmployerID(c echo.Context, id int) ([]*d.Job, error) {
	jobs, err := r.db.SelectJobByEmployerID(id)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *Repo) GetApplications(c echo.Context) ([]*d.Application, error) {
	applications, err := r.db.SelectApplications()
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *Repo) GetApplication(c echo.Context, id int) (*d.Application, error) {
	application, err := r.db.SelectApplicationByID(id)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (r *Repo) GetApplicationsByEmployerID(c echo.Context, employer_id int) ([]*d.Application, error) {
	applications, err := r.db.SelectApplicationsByEmployerID(employer_id)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *Repo) GetApplicationsByTalentID(c echo.Context, talent_id int) ([]*d.Application, error) {
	applications, err := r.db.SelectApplicationsByTalentID(talent_id)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *Repo) CreateApplication(c echo.Context, job_id, talent_id, employer_id int) (*d.Application, error) {
	newApplication := d.NewApplication(job_id, talent_id, employer_id, string(d.Review))

	if err := r.db.CreateApplication(newApplication); err != nil {
		return nil, err
	}

	return newApplication, nil
}

func (r *Repo) UpdateApplicationStatus(c echo.Context, id int, status string) (*d.Application, error) {
	application, err := r.db.SelectApplicationByID(id)
	if err != nil {
		return nil, err
	}

	application.Status = status

	if err := r.db.UpdateApplicationStatus(application); err != nil {
		return nil, err
	}

	return application, nil
}
