package books

type Book struct {
	ID     int
	Title  string `required:"true"`
	Author string `required:"true"`
}
