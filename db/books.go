package db

import (
	"os"

	utils_db "github.com/karlbehrensg/go-clean-architecture-template/utils/db"
)

func SetupBooksDb() *utils_db.PgSetup {
	return &utils_db.PgSetup{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}
