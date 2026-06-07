package storage

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

func SaveToDisk(formFileData []byte) (string, error) {
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create the ./uploads directory: %v", err)
	}

	filename := uuid.New().String() + ".md"
	filepath := path.Join("./uploads", filename)
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create the file at %v", filepath)
	}
	defer file.Close()

	_, err = file.Write(formFileData)
	if err != nil {
		return "", fmt.Errorf("failed to save the form file contents to %v", filepath)
	}

	return filepath, nil
}
