package app

import (
	"File-Processor/config"
	"File-Processor/internal/server"
	"File-Processor/pkg/database"
	"log"
)

// Initialize inicializa la aplicación
func Initialize(config *config.Config) {
	// Conexión a la base de datos
	db, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	// Configuración y puesta en marcha del servidor
	_, err = server.SetupServer(config, db)
	if err != nil {
		log.Fatalf("Error setting up server: %v", err)
	}
}
