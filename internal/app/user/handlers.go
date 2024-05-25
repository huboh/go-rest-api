package user

import (
	"net/http"

	"github.com/huboh/go-rest-api/internal/pkg/json"
)

func handleGetHello(w http.ResponseWriter, r *http.Request) {
	json.Write(w, json.Response{
		Data: "hello",
	})
}
