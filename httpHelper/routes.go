package httpHelper

import (
	"net/http"
)

type Routes struct {
	UrlPrefix   string
	Routes      []*Route
	Middlewares []Middleware
}

func (r *Routes) composeRoutes(mux *http.ServeMux) {
	for _, route := range r.Routes {
		route.UrlPattern = r.UrlPrefix + route.UrlPattern
		mux.HandleFunc(route.UrlPattern, route.composeAction(r.Middlewares...))
	}
}

func RegisterGroupRoutes(mux *http.ServeMux, routes ...*Routes) {
	for _, v := range routes {
		v.composeRoutes(mux)
	}
}
