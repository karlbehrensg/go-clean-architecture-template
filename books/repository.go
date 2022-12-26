package books

import "database/sql"

type repository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) Repository {
	return &repository{client}
}

func (r *repository) CreateBook(book *Book) error {
	err := r.client.QueryRow(`
		INSERT INTO books (title, author)
		VALUES ($1, $2)
		RETURNING id
	`, book.Title, book.Author).Scan(&book.ID)

	if err != nil {
		return err
	}

	return nil
}
