package books

import (
	domain "github.com/karlbehrensg/go-clean-architecture-template/books/domain"
)

type BooksService interface {
	CreateBook(book *domain.Book) (*domain.Book, error)
}
