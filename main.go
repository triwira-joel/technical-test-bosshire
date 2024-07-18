package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/triwira-joel/technical-test-bosshire/config"
	"github.com/triwira-joel/technical-test-bosshire/db"
	h "github.com/triwira-joel/technical-test-bosshire/handler"
	"github.com/triwira-joel/technical-test-bosshire/repo"
	"github.com/triwira-joel/technical-test-bosshire/usecase"
)

func main() {

	cnf := config.Get()
	database := db.ConnectDB(cnf)
	if cnf.InitDB {
		if err := database.InitDB(); err != nil {
			log.Fatal(err)
		}
	}
	repo := repo.NewRepo(database)
	usecase := usecase.NewUseCase(*repo)
	handler := h.NewHttpHandler(*usecase)

	e := echo.New()

	u := e.Group("/users")
	j := e.Group("/jobs")
	a := e.Group("/applications")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	u.GET("/:id", handler.GetUser)
	u.GET("/", handler.GetUsers)
	u.POST("/", handler.CreateUser)

	j.GET("/", handler.GetJobs)
	j.GET("/:id", handler.GetJob)
	j.GET("/:employer_id", handler.GetJobsByEmployerID)
	j.POST("/", handler.CreateJob)

	a.GET("/", handler.GetApplications)
	a.GET("/:id", handler.GetApplication)
	a.GET("/:talent_id", handler.GetApplicationsByTalentID)
	a.GET("/:employer_id", handler.GetJobsByEmployerID)
	a.POST("/", handler.CreateApplication)
	a.PUT("/:id", handler.UpdateApplicationStatus)

	e.Logger.Fatal(e.Start(":1323"))
}
