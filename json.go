package main

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data Response) {
	if data.Code < 100 {
		data.Code = http.StatusOK
	}

	if data.Status == "" {
		if data.Code >= 500 {
			data.Status = StatusError
		} else {
			data.Status = StatusSuccess
		}
	}

	if data.Message == "" {
		data.Message = http.StatusText(data.Code)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(data.Code)
	json.NewEncoder(w).Encode(data)
}
