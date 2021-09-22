package httpHelper

import (
	"fmt"
	"net/http"
)

func RegisterRoute(mux *http.ServeMux, route *Route) {
	mux.HandleFunc(route.UrlPattern, route.composeAction())
}

type Route struct {
	UrlPattern       string
	Method           string
	Action           http.HandlerFunc
	RouteMiddlewares Middlewares
	Options          bool
}

func (r *Route) composeAction(routesMiddlewares ...Middleware) http.HandlerFunc {
	var middlewares []Middleware
	middlewares = append(middlewares, r.RouteMiddlewares...)
	middlewares = append(middlewares, routesMiddlewares...)
	middlewares = append(middlewares, r.checkRequestMethodMiddleware)

	endPointHandlerWithMethod := method(r.Method, r.Action)

	var middlewareChain http.HandlerFunc

	if len(middlewares) == 0 {
		middlewareChain = endPointHandlerWithMethod
	}

	for i := 0; i < len(middlewares); i++ {
		middleware := middlewares[i]
		if i == 0 {
			middlewareChain = middleware(endPointHandlerWithMethod)
		} else {
			middlewareChain = middleware(middlewareChain)
		}
	}

	if r.Options {
		return getOptionsMiddleware(middlewareChain)
	}
	return middlewareChain
}

func (r *Route) checkRequestMethodMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)

			_, err := w.Write([]byte(`{"error":"` + req.Method + ` is not allowed for this route. available: ` + r.Method + `"}`))
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		next.ServeHTTP(w, req)
	}
}
