package auth

import (
	"net/http"

	"github.com/huboh/go-rest-api/internal/app/user"
	"github.com/huboh/go-rest-api/internal/pkg/json"
)

func AuthGuardMiddleware(next http.Handler) http.Handler {
	writeErr := func(w http.ResponseWriter, err error) {
		var msg string

		if err != nil {
			msg = err.Error()
		}

		json.Write(w, json.Response{
			Status:     json.StatusError,
			StatusCode: http.StatusUnauthorized,
			Error: &json.Error{
				Name:    "Unauthorized",
				Message: msg,
			},
		})
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token, err := getAuthHeaderToken(*r)
			if err != nil {
				writeErr(w, err)
				return
			}

			payload, err := tokens.VerifyAccessToken(token)
			if err != nil {
				writeErr(w, err)
				return
			}

			next.ServeHTTP(w, r.WithContext(user.ContextWithUser(r.Context(), payload)))
		},
	)
}
