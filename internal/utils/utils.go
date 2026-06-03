package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err error, status int) error {
	errResp := errorResponse {
		Status: status,
		Message: err.Error(),
	}
	
	data, err := json.Marshal(errResp)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		return err
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