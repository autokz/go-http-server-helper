package httpHelper

import (
	"net/http"
)

type Routes struct {
	UrlPrefix   string
	Routes      []*Route
	Middlewares []Middleware
}

// composeRoutes
// Deprecated: The function is deprecated use the httpHelper2 package.
func (r *Routes) composeRoutes(mux *http.ServeMux) {
	for _, route := range r.Routes {
		route.UrlPattern = r.UrlPrefix + route.UrlPattern
		mux.HandleFunc(route.UrlPattern, route.composeAction(r.Middlewares...))
	}
}

// RegisterGroupRoutes
// Deprecated: The function is deprecated use the httpHelper2 package.
func RegisterGroupRoutes(mux *http.ServeMux, routes ...*Routes) {
	for _, v := range routes {
		v.composeRoutes(mux)
	}
}
