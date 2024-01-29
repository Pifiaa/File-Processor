package routes

import (
	"File-Processor/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes configura las rutas de la aplicación
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	//TODO:Middleware de autenticación

	api.Route("/upload", func(router fiber.Router) {
		router.Post("/", func(c *fiber.Ctx) error {
			return handlers.GetTransactions(db, c)
		})
	})
}
