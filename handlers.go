package main

import "net/http"

var (
	authRouter  = NewRouter("/auth", []Middleware{}, authRoutes)
	usersRouter = NewRouter("/users", []Middleware{}, usersRoutes)
)

var (
	authRoutes  = []Route{}
	usersRoutes = []Route{
		{
			Path:   "/",
			Method: http.MethodGet,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				SendJson(w, Response{
					Data: nil,
				})
			}),
		},
	}
)
