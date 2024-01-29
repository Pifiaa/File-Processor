package app

import (
	"File-Processor/config"
	"File-Processor/internal/server"
	"File-Processor/pkg/database"
	"log"
)

func Initialize(config *config.Config) {
	db, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	_, err = server.SetupServer(config, db)
	if err != nil {
		log.Fatalf("Error setting up server: %v", err)
	}
}
