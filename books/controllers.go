package books

import "github.com/gofiber/fiber/v2"

type controller struct {
	cases Cases
}

func NewController(cases Cases) Controller {
	return &controller{cases}
}

func (c *controller) CreateBook(ctx *fiber.Ctx) error {
	book := &Book{}
	err := ctx.BodyParser(book)
	if err != nil {
		return err
	}

	book, err = c.cases.CreateBook(book)
	if err != nil {
		return err
	}

	return ctx.JSON(book)
}
