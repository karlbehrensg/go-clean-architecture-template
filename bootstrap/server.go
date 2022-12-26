package bootstrap

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/karlbehrensg/go-clean-architecture-template/books"
)

type Server struct {
	App *fiber.App
	Dbs *Dbs
}

func NewServer() *Server {
	server := &Server{}
	server.setup()
	return server
}

func (server *Server) setup() error {
	// Setup dbs
	dbs, err := SetupDbs()
	if err != nil {
		return err
	}
	server.Dbs = dbs

	// Create new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup routes
	books.SetupRoutes(app, server.Dbs.GetDbClient("books"))

	server.App = app

	return nil
}

func (server *Server) Start() error {
	if err := server.setup(); err != nil {
		return err
	}

	// Start server
	return server.App.Listen(":" + os.Getenv("PORT"))
}
