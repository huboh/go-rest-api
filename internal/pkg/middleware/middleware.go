// Package middleware provides functions that add additional functionality to HTTP handlers.
package middleware

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/huboh/go-rest-api/internal/pkg/json"
)

// Middleware is a function that adds additional functionality
// to an http.Handler instance.
type Middleware func(http.Handler) http.Handler

// Logger is a middleware function that logs each incoming HTTP request and
// the time it took to write back a response.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next.ServeHTTP(w, r)

			log.Printf("Request: %s %s took %s\n", r.Method, r.URL.Path, time.Since(now))
		},
	)
}

// GetTimeoutMiddleware is a function that returns a middleware function that
func GetTimeoutMiddleware(d time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx, cancel := context.WithTimeout(
					// get request's original context
					r.Context(),

					// pass the duration
					d,
				)

				// release resources associated with this ctx when func exits
				defer cancel()

				// serve http requests with a new request with context timeout set to d
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}

// PanicRecoverer is a middleware function that recovers from panics in the request handling chain.
//
// If a panic occurs, it captures the panic value, logs an appropriate error message, and sends a JSON response
// with an internal server error status.
func PanicRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if val := recover(); val != nil {
					var msg string

					// TODO(huboh): properly handle `ResponseError` type
					switch val := val.(type) {
					case error:
						msg = val.Error()
					case string:
						msg = val
					default:
						msg = "unknown error"
					}

					json.Write(w, json.Response{
						StatusCode: http.StatusInternalServerError,
						Error: &json.Error{
							Stack:   string(debug.Stack()),
							Message: msg,
						},
					})
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}
