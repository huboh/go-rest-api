package main

import (
	"log"
	"net/http"
	"time"
)

// Middleware is a function that adds additional functionality
// to an http.Handler instance.
type Middleware func(http.Handler) http.Handler

// LoggerMiddleware is a middleware function that logs each incoming HTTP request and
// the time it took to write back a response.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			next.ServeHTTP(w, r)

			log.Printf("Request: %s %s took %s\n", r.Method, r.URL.Path, time.Since(now))
		},
	)
}

// PanicHandlerMiddleware is a middleware function that recovers from panics in the request handling chain.
//
// If a panic occurs, it captures the panic value, logs an appropriate error message, and sends a JSON response
// with an internal server error status.
func PanicHandlerMiddleware(next http.Handler) http.Handler {
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

					SendJson(w, Response{
						StatusCode: http.StatusInternalServerError,
						Error: &ResponseError{
							Message: msg,
						},
					})
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}
