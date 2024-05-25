// Package json provides utilities for working with JSON in HTTP responses.
package json

import (
	jsonEncoder "encoding/json"
	"net/http"
)

// Write writes a JSON response to the provided http.ResponseWriter.
func Write(w http.ResponseWriter, data Response) {
	if data.StatusCode < 100 {
		data.StatusCode = http.StatusOK
	}

	if data.Status == "" {
		if data.StatusCode >= 500 {
			data.Status = StatusError
		} else {
			data.Status = StatusSuccess
		}
	}

	if data.Error != nil {
		data.Status = StatusError
	}

	if data.Message == "" {
		data.Message = http.StatusText(data.StatusCode)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(data.StatusCode)
	jsonEncoder.NewEncoder(w).Encode(data)
}
