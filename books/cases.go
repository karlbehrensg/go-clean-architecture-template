package books

type cases struct {
	repository Repository
}

func NewCases(repository Repository) Cases {
	return &cases{repository}
}

func (c *cases) CreateBook(book *Book) (*Book, error) {
	err := c.repository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
