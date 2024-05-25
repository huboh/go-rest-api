package main

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data Response) {
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
	json.NewEncoder(w).Encode(data)
}
