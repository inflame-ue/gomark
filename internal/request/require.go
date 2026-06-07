package request

import (
	"fmt"
	"net/http"
	"strings"
)

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
