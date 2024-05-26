package auth

import (
	"net/http"

	"github.com/huboh/go-rest-api/internal/pkg/middleware"
	"github.com/huboh/go-rest-api/internal/pkg/router"
)

var (
	Router = router.New(
		// mount path
		RouterPath,

		// middlewares
		[]middleware.Middleware{},

		// routes
		[]router.Route{
			{
				Path:    "/login",
				Method:  http.MethodPost,
				Handler: http.HandlerFunc(handleLogin),
			},
			{
				Path:    "/signup",
				Method:  http.MethodPost,
				Handler: http.HandlerFunc(handleSignup),
			},
			{
				Path:    "/refresh",
				Method:  http.MethodPost,
				Handler: http.HandlerFunc(handleRefresh),
			},
		},
	)
)
