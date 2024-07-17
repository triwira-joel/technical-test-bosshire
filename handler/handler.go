package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/triwira-joel/technical-test-bosshire/domain"
	"github.com/triwira-joel/technical-test-bosshire/usecase"
)

type HTTPHandler struct {
	uc *usecase.UseCase
}

func NewHttpHandler(
	uc usecase.UseCase,
) *HTTPHandler {
	return &HTTPHandler{
		uc: &uc,
	}
}

func (hdl *HTTPHandler) CreateUser(c echo.Context) error {
	createUserReq := new(domain.CreateUserRequest)
	err := json.NewDecoder(c.Request().Body).Decode(createUserReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if createUserReq.Role != "TALENT" && createUserReq.Role != "EMPLOYER" {
		return c.JSON(http.StatusBadRequest, "Role does not exist")
	}

	user, err := hdl.uc.CreateUser(c, createUserReq.Name, createUserReq.Role, createUserReq.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (hdl *HTTPHandler) GetUser(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := hdl.uc.GetUser(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (hdl *HTTPHandler) GetUsers(c echo.Context) error {
	users, err := hdl.uc.GetUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
