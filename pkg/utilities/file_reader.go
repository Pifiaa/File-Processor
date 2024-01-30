package utilities

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
)

// file_reader abre un archivo y devuelve su contenido.
func File_reader(file *multipart.FileHeader) ([]byte, error) {
	// Obtener la extensión del archivo
	ext := filepath.Ext(file.Filename)

	// Verificar si la extensión es .json
	if ext != ".json" {
		return nil, errors.New("el archivo no es de tipo .json")
	}

	// Abrir el archivo
	fileContent, err := file.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()

	// Leer el contenido del archivo
	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}
