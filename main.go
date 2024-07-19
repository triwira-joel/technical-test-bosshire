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
	e.POST("/login", handler.Login)   // LOGIN
	e.POST("/signup", handler.Signup) // SIGNUP

	e.GET("/users/", handler.GetUsers) // TO SHOW ALL USERS

	// employer only
	employer := e.Group("/employer")
	employer.Use(m.JwtEmployerAuthMiddleware())
	employer.GET("/jobs/:employer_id", handler.GetJobsByEmployerID)                     // LIST OF POSTED JOBS
	employer.POST("/jobs", handler.CreateJob)                                           // CREATE JOB POST
	employer.GET("/applications/all/:employer_id", handler.GetApplicationsByEmployerID) // GET APPLICATIONS OF ALL EMPLOYER POSTED JOBS
	employer.PUT("/applications/:id", handler.UpdateApplicationStatus)                  // UPDATE APPLICATION STATUS
	employer.GET("/users/:id", handler.GetUser)                                         // GET USER DETAIL FROM APPLICATION

	// talent only
	talent := e.Group("/talent")
	talent.Use(m.JwtTalentAuthMiddleware())
	talent.GET("/jobs/", handler.GetJobs)                                         // LIST OF ALL JOBS
	talent.GET("/jobs/:id", handler.GetJob)                                       // GET JOB DETAIL
	talent.GET("/applications/:id", handler.GetApplication)                       // GET APPLICATION DETAIL
	talent.GET("/applications/all/:talent_id", handler.GetApplicationsByTalentID) // GET ALL APPLICATIONS FROM THE TALENT ID
	talent.POST("/applications/", handler.CreateApplication)                      // CREATE APPLICATION

	// both party
	both := e.Group("")
	both.Use(m.JwtAuthMiddleware())
	both.GET("/applications/", handler.GetApplications) // TO SHOW ALL APPLICATIONS

	e.Logger.Fatal(e.Start(":1323"))
}
