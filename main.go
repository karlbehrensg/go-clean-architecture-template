// Create server with Fiber
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/karlbehrensg/go-clean-architecture-template/books"
	"github.com/karlbehrensg/go-clean-architecture-template/db"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Create BookDb instance
	booksDb := db.SetupBooksDb()
	booksDb.Connect()
	if err != nil {
		panic(err)
	}
	defer booksDb.Close()

	// Create new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup routes
	books.SetupRoutes(app, booksDb.Client)

	// Start server
	app.Listen(":3000")
}
