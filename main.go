package main

import (
	"fmt"
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
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:id", handler.GetUser)
	e.GET("/user", handler.GetUsers)
	e.POST("/user", handler.CreateUser)

	fmt.Println(database)
	e.Logger.Fatal(e.Start(":1323"))
}
