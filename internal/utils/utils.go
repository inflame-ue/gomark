package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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

	data, err := json.Marshal(errResp)
	if err != nil {
		return fmt.Errorf("failed to marhal the error: %v", err)
	}

	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write the data: %v", err)
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
