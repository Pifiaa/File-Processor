package main

import (
	"File-Processor/config"
	"File-Processor/internal/app"
	"log"
)

func main() {
	// Obtiene la configuración de la aplicación
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	// Crear una nueva instancia de la aplicación y la inicializa
	app.Initialize(config)

	//TODO: Documentación del programa y la api
}
