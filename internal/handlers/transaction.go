package handlers

import (
	"File-Processor/internal/models"
	"File-Processor/pkg/utilities"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	// Número de workers para procesamiento de gouroutines
	numWorkers int = 10

	// Tamaño del lote para procesamiento
	batchSize int = 150
)

func GetTransactions(db *gorm.DB, c *fiber.Ctx) error {
	FormFile, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := utilities.File_reader(FormFile)
	if err != nil {
		log.Fatal(err)
	}

	var transactions []models.Invoicings
	err = json.Unmarshal(fileContent, &transactions)
	if err != nil {
		log.Fatal(err)
	}

	processTransactions(transactions, db)

	fmt.Print("Hola mundo")

	return nil
}

func processTransactions(transactions []models.Invoicings, db *gorm.DB) {
	var wg sync.WaitGroup
	errChan := make(chan error)
	stopChan := make(chan struct{}) // Canal para detener la ejecución de goroutines

	for i := 0; i < len(transactions); i += batchSize {
		end := i + batchSize
		if end > len(transactions) {
			end = len(transactions)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				// Procesar la transacción
				transaction := transactions[j]
				// log.Printf("\n\nTransacción: %+v\n", transaction)

				if err := db.Create(&transaction).Error; err != nil {
					log.Printf("Error al insertar la transacción en la base de datos: %v", err)
					errChan <- err
					// Enviar señal para detener otras goroutines
					stopChan <- struct{}{}
					return
				}
				//TODO: Operación sobre los datos
			}
		}(i, end)
	}

	// Monitorear errores
	go func() {
		for err := range errChan {
			// Manejar el error como desees
			log.Printf("Error encontrado: %v", err)
		}
		// Cerrar el canal de detención cuando no hay más errores
		close(stopChan)
	}()

	go func() {
		// Esperar a que todas las goroutines terminen
		wg.Wait()
		// Cerrar el canal de errores después de que todas las goroutines terminen
		close(errChan)
	}()

	// Esperar hasta que se cierre el canal de detención
	<-stopChan
}
