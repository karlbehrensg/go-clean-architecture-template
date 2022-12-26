// Tests for repository.go
package books

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("should return error if query fails", func(t *testing.T) {
		book := &Book{
			Title:  "The Hitchhiker's Guide to the Galaxy",
			Author: "Douglas Adams",
		}

		mock.ExpectQuery("INSERT INTO books").
			WithArgs(book.Title, book.Author).
			WillReturnError(err)

		repo := NewRepository(db)
		err = repo.CreateBook(book)
		if err == nil {
			t.Errorf("error was expected while creating book")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("should return no error if query succeeds", func(t *testing.T) {
		book := &Book{
			Title:  "The Hitchhiker's Guide to the Galaxy",
			Author: "Douglas Adams",
		}

		mock.ExpectQuery("INSERT INTO books").
			WithArgs(book.Title, book.Author).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		repo := NewRepository(db)
		err = repo.CreateBook(book)
		if err != nil {
			t.Errorf("error was not expected while creating book")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
