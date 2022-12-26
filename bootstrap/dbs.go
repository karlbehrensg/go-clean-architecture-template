package bootstrap

import (
	"database/sql"

	"github.com/karlbehrensg/go-clean-architecture-template/db"
	utils_db "github.com/karlbehrensg/go-clean-architecture-template/utils/db"
)

type Db struct {
	Name    string
	PgSetup *utils_db.PgSetup
}

type Dbs struct {
	Connections []Db
}

func SetupDbs() (*Dbs, error) {
	dbs := &Dbs{}

	// Create BookDb instance
	booksDb := db.SetupBooksDb()
	if err := booksDb.Connect(); err != nil {
		return nil, err
	}

	// Add BookDb to Dbs
	dbs.Connections = append(dbs.Connections, Db{Name: "books", PgSetup: booksDb})

	return dbs, nil
}

func (dbs *Dbs) Close() {
	for _, db := range dbs.Connections {
		db.PgSetup.Close()
	}
}

func (dbs *Dbs) GetDbClient(name string) *sql.DB {
	for _, db := range dbs.Connections {
		if db.Name == name {
			return db.PgSetup.Client
		}
	}
	return nil
}
