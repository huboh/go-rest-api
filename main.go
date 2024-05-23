package main

import (
	"log"
	"net/http"
	"os"

	env "github.com/joho/godotenv"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	if err := env.Load(); err != nil {
		return err
	}

	var (
		host   = os.Getenv("HOST")
		port   = os.Getenv("PORT")
		router = NewRouter("/", getMiddlewares(), getRoutes())
		server = NewServer(
			NewServerConfig(
				host,
				port,
				router,
			),
		)
	)

	defer server.Close()

	log.Println("server listening on", server.httpSvr.Addr)

	if err := server.Listen(); err != nil {
		return err
	}

	return nil
}

func getRoutes() []Route {
	return []Route{
		{
			Path:   "/healthz",
			Method: http.MethodGet,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				SendJson(w, Response{
					Data: "active",
				})
			}),
		},
		{
			Path:    authRouter.Prefix,
			Handler: authRouter,
		},
		{
			Path:    usersRouter.Prefix,
			Handler: usersRouter,
		},
	}
}

func getMiddlewares() []Middleware {
	return []Middleware{}
}
