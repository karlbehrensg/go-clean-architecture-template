package books

import (
	"testing"

	domain "github.com/karlbehrensg/go-clean-architecture-template/books/domain"
)

// ***************** Mocks *****************

type MockBooksRepository struct {
}

func (r *MockBooksRepository) CreateBook(book *domain.Book) (*domain.Book, error) {
	return book, nil
}

// ***************** Tests *****************

func TestBooksService_CreateBook(t *testing.T) {
	type response struct {
		book *domain.Book
		err  error
	}

	booksRepository := &MockBooksRepository{}
	booksService := NewBooksService(booksRepository)

	book1 := &domain.Book{
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
	}

	book2 := &domain.Book{
		Title: "Foundation",
	}

	book3 := &domain.Book{
		Author: "Isaac Asimov",
	}

	var testCases = []struct {
		name string
		book *domain.Book
		want response
	}{
		{"Create book", book1, response{book1, nil}},
		{"Try to create book with missing author", book2, response{nil, errMissingAuthor}},
		{"Try to create book with missing title", book3, response{nil, errMissingTitle}},
	}

	for _, tc := range testCases {
		tc := tc // rebind tc to avoid bugs with parallel tests
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := booksService.CreateBook(tc.book)

			if err != tc.want.err {
				t.Errorf("got error %v, want %v", err, tc.want.err)
			}

			if got != tc.want.book {
				t.Errorf("got %v, want %v", got, tc.want.book)
			}

		})
	}
}
