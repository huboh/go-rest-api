package user

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
				Path:    "/hello",
				Method:  http.MethodGet,
				Handler: http.HandlerFunc(handleGetHello),
			},
		},
	)
)
