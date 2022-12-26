package books

import (
	"database/sql"
)

func SetupController(db *sql.DB) Controller {
	repository := NewRepository(db)
	cases := NewCases(repository)
	return NewController(cases)
}
