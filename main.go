// Create server with Fiber
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/karlbehrensg/go-clean-architecture-template/books"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup routes
	books.SetupRoutes(app, db)

	// Start server
	app.Listen(":3000")
}
