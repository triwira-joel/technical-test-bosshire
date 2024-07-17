package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
	InitDB   bool
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Get() *Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("error when load env %s", err.Error())
	}

	init, err := strconv.ParseBool(os.Getenv("INIT_DB"))
	if err != nil {
		log.Fatalf("error when load env %s", err.Error())
	}

	return &Config{
		Database: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		InitDB: init,
	}
}
