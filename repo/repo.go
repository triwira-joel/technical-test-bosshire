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

func (r *Repo) CreateUser(c echo.Context, name string, role string, email string) (*d.User, error) {
	newUser := d.NewUser(name, role, email)

	if err := r.db.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *Repo) GetUsers(c echo.Context) ([]*d.User, error) {
	users, err := r.db.SelectUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
