package server

import (
	"File-Processor/config"
	"File-Processor/internal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupServer(config *config.Config, db *gorm.DB) (*fiber.App, error) {
	// Inicia la instancia de Fiber
	app := fiber.New()

	// Verifica si la instancia de Fiber se creó correctamente
	if app == nil {
		return nil, fmt.Errorf("error al crear la instancia de Fiber")
	}

	routes.SetupRoutes(app, db, config)

	port := fmt.Sprintf(":%s", config.Server.Port)

	// Inicia el servidor
	err := app.Listen(port)
	if err != nil {
		return nil, fmt.Errorf("error al iniciar el servidor: %v", err)
	}

	return app, nil
}
