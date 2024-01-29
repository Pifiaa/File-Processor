package database

import (
	"File-Processor/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect establece la conexión a la base de datos .
func Connect(config *config.Config) (*gorm.DB, error) {
	// Construir el DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	// Abrir la conexión a la base de datos utilizando gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	return db, nil
}
