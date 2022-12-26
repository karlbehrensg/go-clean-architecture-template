package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PgSetup struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Client   *sql.DB
}

func (pg *PgSetup) Connect() error {
	client, err := sql.Open("postgres", pg.getConnString())
	if err != nil {
		log.Fatal(err)
		return err
	}
	pg.Client = client
	return nil
}

func (pg *PgSetup) Close() {
	pg.Client.Close()
}

func (pg *PgSetup) getConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Dbname,
	)
}
