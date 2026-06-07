package respond

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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

func WriteHTML(w http.ResponseWriter, html []byte) error {
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write(html)
	if err != nil {
		return err
	}

	return nil
}
