package usecase

import (
	"github.com/labstack/echo/v4"
	d "github.com/triwira-joel/technical-test-bosshire/domain"
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

func (u *UseCase) CreateUser(c echo.Context, name string, role string, email string) (*d.User, error) {
	user, err := u.repo.CreateUser(c, name, role, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UseCase) GetUsers(c echo.Context) ([]*d.User, error) {
	users, err := u.repo.GetUsers(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}
