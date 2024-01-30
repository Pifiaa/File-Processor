package routes

import (
	"File-Processor/config"
	"File-Processor/internal/handlers"
	"File-Processor/internal/routes/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes configura las rutas de la aplicación
func SetupRoutes(app *fiber.App, db *gorm.DB, config *config.Config) {
	api := app.Group("/api")

	jwt := middlewares.NewAuthMiddleware(config.Token.Key)

	api.Route("/login", func(router fiber.Router) {
		router.Post("/", func(c *fiber.Ctx) error {
			return handlers.AuthenticateUser(c, db, config)
		})
	})

	api.Route("/upload", func(router fiber.Router) {
		router.Use(jwt)
		router.Post("/", func(c *fiber.Ctx) error {
			return handlers.GetTransactions(db, c)
		})
	})
}
