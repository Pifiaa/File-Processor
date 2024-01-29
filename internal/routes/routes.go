package routes

import (
	"File-Processor/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Route("/upload", func(router fiber.Router) {
		router.Post("/", func(c *fiber.Ctx) error {
			return handlers.GetTransactions(db, c)
		})
	})
}
