package auth

import (
	"net/http"

	"github.com/huboh/go-rest-api/internal/pkg/json"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	resp, err := login()

	if err != nil {
		json.Write(w, json.Response{
			Error: &json.Error{
				Message: err.Error(),
			},
		})
		return
	}

	json.Write(w, json.Response{
		Data: resp,
	})
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	resp, err := signUp()

	if err != nil {
		json.Write(w, json.Response{
			Error: &json.Error{
				Message: err.Error(),
			},
		})
		return
	}

	json.Write(w, json.Response{
		Data: resp,
	})
}
