package books

import "github.com/gofiber/fiber/v2"

type Repository interface {
	CreateBook(book *Book) error
}

type Cases interface {
	CreateBook(book *Book) (*Book, error)
}

type Controller interface {
	CreateBook(ctx *fiber.Ctx) error
}
