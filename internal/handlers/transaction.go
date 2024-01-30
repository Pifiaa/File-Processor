package handlers

import (
	"File-Processor/internal/models"
	"File-Processor/pkg/utilities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// Response estructura para las respuestas
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// GetTransactions maneja la carga y procesamiento de transacciones
func GetTransactions(db *gorm.DB, c *fiber.Ctx) error {
	//Obtiene el archivo
	FormFile, err := c.FormFile("file")
	if err != nil {
		return handleError(c, http.StatusBadRequest, fmt.Sprintf("Error al obtener el archivo: %s", err.Error()))
	}

	//Leer contenido del archivo
	fileContent, err := utilities.File_reader(FormFile)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, fmt.Sprintf("Error al leer el archivo: %s", err.Error()))
	}

	//Decodifica contenido del archivo
	var transactions []models.Invoicings
	err = json.Unmarshal(fileContent, &transactions)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, fmt.Sprintf("Error al decodificar el archivo JSON: %s", err.Error()))
	}

	//Procesa las transacciones
	err = processTransactions(transactions, db)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, fmt.Sprintf("Error al procesar las transacciones: %s", err.Error()))
	}

	return c.JSON(Response{Message: "Transacciones procesadas correctamente"})
}

func processTransactions(transactions []models.Invoicings, db *gorm.DB) error {
	//Tamaño del lote para procesamiento
	batchSize := 700
	//Calcula cantidad de lotes
	batchCount := (len(transactions) + batchSize - 1) / batchSize

	var transactionGroup errgroup.Group

	//Procesa cada lote
	for i := 0; i < batchCount; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > len(transactions) {
			end = len(transactions)
		}

		transactionGroup.Go(func() error {
			//Inicia transacción de base de datos
			databaseTransaction := db.Begin()
			defer func() {
				if r := recover(); r != nil {
					databaseTransaction.Rollback()
				}
			}()

			//Inserta transacciones
			for _, transaction := range transactions[start:end] {
				if err := databaseTransaction.Create(&transaction).Error; err != nil {
					databaseTransaction.Rollback()
					return fmt.Errorf("error al insertar la transacción en la base de datos: %w", err)
				}
				// TODO: Operación sobre los datos
				//TODO: Enviar datos a la API de Konecta
			}
			return databaseTransaction.Commit().Error
		})
	}

	//Espera a que se completen todas las transacciones
	if err := transactionGroup.Wait(); err != nil {
		return err
	}

	return nil
}

// handleError maneja los errores y envía una respuesta JSON
func handleError(c *fiber.Ctx, status int, errorMessage string) error {
	return c.Status(status).JSON(Response{Error: errorMessage})
}
