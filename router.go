package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Middleware
type Middleware func(http.Handler) http.Handler

// Route represents an api route
type Route struct {
	// Path is request Path
	Path string

	// Method is the route http request Method
	Method string

	// Handler is the route http request Handler or a router to handle request received to `path`
	Handler http.Handler
}

func NewRoute(m string, p string, h http.Handler) Route {
	return Route{
		Path:    p,
		Method:  m,
		Handler: h,
	}
}

// Router
type Router struct {
	// mux is the underlying multiplexer
	mux *http.ServeMux

	// Prefix is the router path prefix
	Prefix string

	// Routes
	Routes []Route

	// Middlewares
	Middlewares []Middleware
}

func NewRouter(prefix string, mws []Middleware, routes []Route) *Router {
	router := &Router{
		mux:         new(http.ServeMux),
		Prefix:      prefix,
		Routes:      routes,
		Middlewares: mws,
	}

	// register routes
	router.registerRoutes()

	return router
}

// ServeHTTP makes Router implements http.Handler interface
func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(writer, req)
}

// registerRoutes registers the router routes with the specified handlers
func (r *Router) registerRoutes() {
	log.Printf("new router at %s\n", r.Prefix)

	for _, route := range r.Routes {
		var (
			path     = Must(url.JoinPath(r.Prefix, route.Path))
			handler  = route.Handler
			fullPath = strings.TrimSpace(fmt.Sprintf("%s %s", route.Method, path))
		)

		if !strings.HasSuffix(fullPath, "/") {
			fullPath += "/"
		}

		// map to path handler
		r.mux.Handle(fullPath, handler)

		switch handler.(type) {
		case *Router:

		case http.Handler, http.HandlerFunc:
			log.Printf("new handler mapped for %s\n", fullPath)
		}
	}
}
