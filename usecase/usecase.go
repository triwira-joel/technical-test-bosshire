package usecase

import (
	"github.com/labstack/echo/v4"
	d "github.com/triwira-joel/technical-test-bosshire/domain"
	"github.com/triwira-joel/technical-test-bosshire/helper"
	repo "github.com/triwira-joel/technical-test-bosshire/repo"
)

type UseCase struct {
	repo *repo.Repo
}

func NewUseCase(
	repo repo.Repo,
) *UseCase {
	return &UseCase{
		repo: &repo,
	}
}

func (u *UseCase) GetUser(c echo.Context, id int) (*d.User, error) {
	user, err := u.repo.GetUser(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UseCase) CreateUser(c echo.Context, name string, role string, email string, password string) (string, error) {
	user, err := u.repo.CreateUser(c, name, role, email, password)
	if err != nil {
		return "", err
	}

	jwt, err := helper.CreateJWTToken(user)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (u *UseCase) Login(c echo.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmailAndPassword(c, email, password)
	if err != nil {
		return "", err
	}

	jwt, err := helper.CreateJWTToken(user)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (u *UseCase) GetUsers(c echo.Context) ([]*d.User, error) {
	users, err := u.repo.GetUsers(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UseCase) GetJobs(c echo.Context) ([]*d.Job, error) {
	jobs, err := u.repo.GetJobs(c)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (u *UseCase) CreateJob(c echo.Context, name string, description string, employer_id int) (*d.Job, error) {
	job, err := u.repo.CreateJob(c, name, description, employer_id)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (u *UseCase) GetJob(c echo.Context, id int) (*d.Job, error) {
	job, err := u.repo.GetJob(c, id)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (u *UseCase) GetJobsByEmployerID(c echo.Context, id int) ([]*d.Job, error) {
	jobs, err := u.repo.GetJobsByEmployerID(c, id)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (u *UseCase) GetApplications(c echo.Context) ([]*d.Application, error) {
	applications, err := u.repo.GetApplications(c)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (u *UseCase) GetApplication(c echo.Context, id int) (*d.Application, error) {
	application, err := u.repo.GetApplication(c, id)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (u *UseCase) GetApplicationsByEmployerID(c echo.Context, employer_id int) ([]*d.Application, error) {
	applications, err := u.repo.GetApplicationsByEmployerID(c, employer_id)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (u *UseCase) GetApplicationsByTalentID(c echo.Context, talent_id int) ([]*d.Application, error) {
	applications, err := u.repo.GetApplicationsByTalentID(c, talent_id)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (u *UseCase) CreateApplication(c echo.Context, job_id, talent_id, employer_id int) (*d.Application, error) {
	application, err := u.repo.CreateApplication(c, job_id, talent_id, employer_id)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (u *UseCase) UpdateApplicationStatus(c echo.Context, id int, status string) (*d.Application, error) {
	application, err := u.repo.UpdateApplicationStatus(c, id, status)
	if err != nil {
		return nil, err
	}
	return application, nil
}
