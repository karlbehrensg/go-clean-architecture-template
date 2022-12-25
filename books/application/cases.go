package books

import (
	"errors"

	domain "github.com/karlbehrensg/go-clean-architecture-template/books/domain"
)

type booksService struct {
	booksRepository domain.BooksRepository
}

var (
	errMissingAuthor = errors.New("missing author")
	errMissingTitle  = errors.New("missing title")
)

func NewBooksService(booksRepository domain.BooksRepository) BooksService {
	return &booksService{
		booksRepository,
	}
}

func (r *booksService) CreateBook(book *domain.Book) (*domain.Book, error) {
	return r.booksRepository.CreateBook(book)
}
