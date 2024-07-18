package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/triwira-joel/technical-test-bosshire/config"
	"github.com/triwira-joel/technical-test-bosshire/db"
	h "github.com/triwira-joel/technical-test-bosshire/handler"
	m "github.com/triwira-joel/technical-test-bosshire/middleware"
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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// public
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", handler.Login)
	e.POST("/signup", handler.Signup)

	e.GET("/users/", handler.GetUsers)

	// employer only
	employer := e.Group("/employer")
	employer.Use(m.JwtEmployerAuthMiddleware())
	employer.GET("/jobs/:employer_id", handler.GetJobsByEmployerID)
	employer.POST("/jobs", handler.CreateJob)
	employer.GET("/applications/:employer_id", handler.GetApplicationsByEmployerID)
	employer.PUT("/applications/:id", handler.UpdateApplicationStatus)
	employer.GET("/users/:id", handler.GetUser)

	// talent only
	talent := e.Group("/talent")
	talent.Use(m.JwtTalentAuthMiddleware())
	talent.GET("/jobs/", handler.GetJobs)
	talent.GET("/jobs/:id", handler.GetJob)
	talent.GET("/applications/:id", handler.GetApplication)
	talent.GET("/applications/:talent_id", handler.GetApplicationsByTalentID)
	talent.POST("/applications/", handler.CreateApplication)

	// both party
	both := e.Group("")
	both.Use(m.JwtAuthMiddleware())
	both.GET("/applications/", handler.GetApplications)

	e.Logger.Fatal(e.Start(":1323"))
}
