package auth

import (
	"net/http"

	"github.com/huboh/go-rest-api/internal/pkg/json"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		creds loginCredentials
	)

	err = json.UnmarshalBody(r, &creds)

	if err != nil {
		json.Write(w, json.Response{
			StatusCode: http.StatusBadRequest,
			Error:      json.ErrorFromErr(err, "", ""),
		})
		return
	}

	result, err := login(r.Context(), creds)

	if err != nil {
		json.Write(w, json.Response{
			StatusCode: http.StatusBadRequest,
			Error:      json.ErrorFromErr(err, "", ""),
		})
		return
	}

	json.Write(w, json.Response{
		Data: result,
	})
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	result, err := signUp(r.Context())

	if err != nil {
		json.Write(w, json.Response{
			StatusCode: http.StatusBadRequest,
			Error:      json.ErrorFromErr(err, "", ""),
		})
		return
	}

	json.Write(w, json.Response{
		Data: result,
	})
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	result, err := refresh(r.Context())

	if err != nil {
		json.Write(w, json.Response{
			StatusCode: http.StatusBadRequest,
			Error:      json.ErrorFromErr(err, "", ""),
		})
		return
	}

	json.Write(w, json.Response{
		Data: result,
	})
}
