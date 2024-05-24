package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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
	r := &Router{
		mux:         new(http.ServeMux),
		Prefix:      prefix,
		Routes:      routes,
		Middlewares: mws,
	}

	// register routes
	r.registerRoutes()

	return r
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

		if _, ok := handler.(*Router); ok && !strings.HasSuffix(fullPath, "/") {
			fullPath += "/"
		}

		// map to path handler
		r.mux.Handle(fullPath, r.registerMiddlewares(handler))

		switch handler.(type) {
		case *Router:

		case http.Handler, http.HandlerFunc:
			log.Printf("new handler mapped for %s\n", fullPath)
		}
	}
}

// registerMiddlewares registers the router middlewares
func (r *Router) registerMiddlewares(h http.Handler) http.Handler {
	for _, middleware := range r.Middlewares {
		h = middleware(h)
	}

	return h
}
