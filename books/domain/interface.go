package books

type BooksRepository interface {
	CreateBook(book *Book) (*Book, error)
}
