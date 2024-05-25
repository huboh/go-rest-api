package router

import "net/http"

// Route represents an api route
type Route struct {
	// Tag is used to add metadata to the route
	Tag []string

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
