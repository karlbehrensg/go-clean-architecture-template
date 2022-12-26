package books

import (
	"errors"
	"testing"
)

type mockRepository struct{}

func newMockRepository() Repository {
	return &mockRepository{}
}

func (m *mockRepository) CreateBook(book *Book) error {
	if book.ID == 1 {
		return errors.New("error")
	}

	book.ID = 1
	return nil
}

func TestCasesNewCases(t *testing.T) {
	repository := newMockRepository()
	cases := NewCases(repository)

	if cases == nil {
		t.Errorf("cases was not expected to be nil")
	}
}

func TestCasesCreateBook(t *testing.T) {
	repository := newMockRepository()
	t.Run("Return error because repository return error", func(t *testing.T) {
		cases := NewCases(repository)

		book := &Book{
			ID:     1,
			Title:  "The Hitchhiker's Guide to the Galaxy",
			Author: "Douglas Adams",
		}

		_, err := cases.CreateBook(book)
		if err == nil {
			t.Errorf("error was expected while creating book")
		}
	})

	t.Run("Return no error because repository return no error", func(t *testing.T) {
		cases := NewCases(repository)

		book := &Book{
			Title:  "The Hitchhiker's Guide to the Galaxy",
			Author: "Douglas Adams",
		}

		_, err := cases.CreateBook(book)
		if err != nil {
			t.Errorf("error was not expected while creating book")
		}
	})
}
