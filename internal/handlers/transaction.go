package handlers

import (
	"File-Processor/pkg/utilities"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	fmt.Print(fileContent)

	//transactions := utilities.Decode_json(fileContent)

	/*// Iniciar el contador
	startTime := time.Now()

	// Abrir el archivo JSON
	FormFile, err := c.FormFile("file")
	file, err := FormFile.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decodificar el archivo JSON
	var transactions []Transaction
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&transactions); err != nil {
		log.Fatal(err)
	}

	// Procesar el lote de transacciones en goroutines
	var wg sync.WaitGroup
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
				// transaction := transactions[j]
				// log.Printf("\n\nTransacción: %+v\n", transaction)

				//TODO: Operación sobre los datos
			}
		}(i, end)
	}
	wg.Wait()

	// Calcula el tiempo transcurrido
	elapsedTime := time.Since(startTime)
	log.Printf("Tiempo transcurrido: %v", elapsedTime)

	return nil*/

	fmt.Print("Hola mundo")

	return nil
}
