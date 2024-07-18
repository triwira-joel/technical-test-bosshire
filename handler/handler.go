package handler

import (
	"encoding/json"
	"fmt"
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
		return c.JSON(http.StatusBadRequest, err.Error()) // TODO tambahin error message
	}

	if createUserReq.Role != "TALENT" && createUserReq.Role != "EMPLOYER" {
		return c.JSON(http.StatusBadRequest, "Role does not exist")
	}

	user, err := hdl.uc.CreateUser(c, createUserReq.Name, createUserReq.Role, createUserReq.Email, createUserReq.Password)
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

func (hdl *HTTPHandler) GetJobs(c echo.Context) error {
	jobs, err := hdl.uc.GetJobs(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jobs)
}

func (hdl *HTTPHandler) CreateJob(c echo.Context) error {
	createJobReq := new(domain.CreateJobRequest)
	err := json.NewDecoder(c.Request().Body).Decode(createJobReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	job, err := hdl.uc.CreateJob(c, createJobReq.Name, createJobReq.Description, createJobReq.EmployerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, job)
}

func (hdl *HTTPHandler) GetJob(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	job, err := hdl.uc.GetJob(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, job)
}

func (hdl *HTTPHandler) GetJobsByEmployerID(c echo.Context) error {
	idstr := c.Param("employer_id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(id)

	jobs, err := hdl.uc.GetJobsByEmployerID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jobs)
}

func (hdl *HTTPHandler) GetApplications(c echo.Context) error {
	applications, err := hdl.uc.GetApplications(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, applications)
}

func (hdl *HTTPHandler) GetApplication(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(id)

	application, err := hdl.uc.GetApplication(c, id)
	fmt.Println("-- HANDLER MASUK --")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, application)
}

func (hdl *HTTPHandler) GetApplicationsByEmployerID(c echo.Context) error {
	idstr := c.Param("employer_id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(id)

	applications, err := hdl.uc.GetApplicationsByEmployerID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, applications)
}

func (hdl *HTTPHandler) GetApplicationsByTalentID(c echo.Context) error {
	idstr := c.Param("talent_id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(id)

	applications, err := hdl.uc.GetApplicationsByTalentID(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, applications)
}

func (hdl *HTTPHandler) CreateApplication(c echo.Context) error {
	createApplicationReq := new(domain.CreateApplicationRequest)
	err := json.NewDecoder(c.Request().Body).Decode(createApplicationReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	application, err := hdl.uc.CreateApplication(c, createApplicationReq.JobID, createApplicationReq.TalentID, createApplicationReq.EmployerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}

func (hdl *HTTPHandler) UpdateApplicationStatus(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(id)

	updateApplicationReq := new(domain.UpdateApplicationRequest)
	err = json.NewDecoder(c.Request().Body).Decode(updateApplicationReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if updateApplicationReq.Status != string(domain.Interview) && updateApplicationReq.Status != string(domain.Accept) && updateApplicationReq.Status != string(domain.Reject) {
		return c.JSON(http.StatusBadRequest, "Invalid Status")
	}

	application, err := hdl.uc.UpdateApplicationStatus(c, id, updateApplicationReq.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, application)
}
