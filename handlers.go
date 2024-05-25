package main

import (
	"net/http"
	"time"

	"github.com/huboh/go-rest-api/internal/pkg/json"
)

func handleGetHealthz(w http.ResponseWriter, r *http.Request) {
	// panic("hello")
	json.Write(w, json.Response{
		Data: time.Now().Format(time.RFC3339),
	})
}
