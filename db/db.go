package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/triwira-joel/technical-test-bosshire/config"
	"github.com/triwira-joel/technical-test-bosshire/constants"
)

type DBHandler struct {
	DB *sql.DB
}

func (s *DBHandler) InitDB() error {
	return s.CreateUserTable()
}

func ConnectDB(cnf *config.Config) *DBHandler {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(constants.ConnectDBFail + " | " + err.Error())
	}

	return &DBHandler{
		DB: db,
	}
}

func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		log.Println(constants.ClosingDBReadFailed + " | " + err.Error())
	} else {
		log.Println(constants.ClosingDBReadSuccess)
	}
}
