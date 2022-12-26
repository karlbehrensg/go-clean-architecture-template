package books

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	controller := SetupController(db)

	routes := app.Group("/books")

	routes.Post("/", controller.CreateBook)
}
