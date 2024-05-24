package main

import "net/http"

var (
	authRouter  = NewRouter("/auth", []Middleware{}, authRoutes)
	usersRouter = NewRouter("/users", []Middleware{}, usersRoutes)
)

var (
	authRoutes = []Route{
		{
			Path:   "/login",
			Method: http.MethodPost,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp, err := Login()

				if err != nil {
					SendJson(w, Response{
						Error: &ResponseError{
							Message: err.Error(),
						},
					})
					return
				}

				SendJson(w, Response{
					Data: resp,
				})
			}),
		},
		{
			Path:   "/signup",
			Method: http.MethodPost,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp, err := SignUp()

				if err != nil {
					SendJson(w, Response{
						Error: &ResponseError{
							Message: err.Error(),
						},
					})
					return
				}

				SendJson(w, Response{
					Data: resp,
				})
			}),
		},
	}

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
