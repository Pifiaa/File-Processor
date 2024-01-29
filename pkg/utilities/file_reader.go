package utilities

import (
	"io"
	"log"
	"mime/multipart"
)

// file_reader abre un archivo y devuelve su contenido.
func File_reader(file *multipart.FileHeader) ([]byte, error) {
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
