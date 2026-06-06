package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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

func RequireMethod(r *http.Request, method string) error {
	if r.Method != strings.ToUpper(method) {
		return fmt.Errorf("expected the HTTP method to be %v, got %v instead", method, r.Method)
	}
	return nil
}

func RequireContentType(r *http.Request, contentType string) error {
	reqContentType := r.Header.Get("Content-Type")
	if reqContentType != strings.ToLower(contentType) {
		return fmt.Errorf("expected the content type to be %v, got %v instead",
			contentType, reqContentType)
	}
	return nil
}

func WriteErrorAndLog(w http.ResponseWriter, err error, status int) error {
	log.Print(err)

	errResp := errorResponse{
		Status:  status,
		Message: err.Error(),
	}

	data, marshalErr := json.Marshal(errResp)
	if marshalErr != nil {
		return fmt.Errorf("failed to marhal the error: %v", marshalErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, writeErr := w.Write(data)
	if writeErr != nil {
		return fmt.Errorf("failed to write the data: %v", writeErr)
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
