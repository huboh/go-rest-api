package main

import (
	"log"
	"net/http"
	"os"

	"github.com/huboh/go-rest-api/internal/app/auth"
	"github.com/huboh/go-rest-api/internal/app/user"

	"github.com/huboh/go-rest-api/internal/pkg/env"
	"github.com/huboh/go-rest-api/internal/pkg/middleware"
	"github.com/huboh/go-rest-api/internal/pkg/router"
	"github.com/huboh/go-rest-api/internal/pkg/server"
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

	server := server.New(
		server.NewConfig(
			// host
			os.Getenv("HOST"),

			// port
			os.Getenv("PORT"),

			// router
			router.New("/", getMiddlewares(), getRoutes()),
		),
	)

	defer server.Stop()

	log.Println("server listening on", server.Configs.Addr)

	if err := server.Start(); err != nil {
		return err
	}

	return nil
}

func getRoutes() []router.Route {
	return []router.Route{
		{
			Path:    user.RouterPath,
			Handler: user.Router,
		},
		{
			Path:    auth.RouterPath,
			Handler: auth.Router,
		},
		{
			Path:    "/healthz",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(handleGetHealthz),
		},
	}
}

func getMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{
		middleware.PanicRecoverer,
		middleware.Logger,
	}
}
