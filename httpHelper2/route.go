package httpHelper2

import (
	"net/http"
)

// Route represents an individual route in the router.
type Route struct {
	pattern     string
	method      string
	handler     func(http.ResponseWriter, *http.Request)
	original    func(http.ResponseWriter, *http.Request)
	middlewares []Middleware
}

// Middleware appends one or more middleware functions to the route.
func (r *Route) Middleware(middlewares ...Middleware) *Route {
	r.middlewares = append(r.middlewares, middlewares...)
	r.handler = r.original
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		r.handler = r.middlewares[i](r.handler)
	}
	return r
}

// methodMiddleware checks if the HTTP request method matches the route's method.
func (r *Route) methodMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		if request.Method != r.method {
			writeJson(rw, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return
		}
		next.ServeHTTP(rw, request)
	}
}

// Get registers a new route with the GET method.
func (r *Router) Get(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodGet, pattern, handler)
}

// Head registers a new route with the Head method.
func (r *Router) Head(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodHead, pattern, handler)
}

// Post registers a new route with the Post method.
func (r *Router) Post(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodPost, pattern, handler)
}

// Put registers a new route with the Put method.
func (r *Router) Put(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodPut, pattern, handler)
}

// Patch registers a new route with the Patch method.
func (r *Router) Patch(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodPatch, pattern, handler)
}

// Delete registers a new route with the Delete method.
func (r *Router) Delete(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodDelete, pattern, handler)
}

// Connect registers a new route with the Connect method.
func (r *Router) Connect(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodConnect, pattern, handler)
}

// Options registers a new route with the Options method.
func (r *Router) Options(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodOptions, pattern, handler)
}

// Trace registers a new route with the Trace method.
func (r *Router) Trace(pattern string, handler http.HandlerFunc) *Route {
	return r.registerRoute(http.MethodTrace, pattern, handler)
}

// HandleFunc registers a new route with a custom HTTP method.
func (r *Router) HandleFunc(method string, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{
		pattern:     pattern,
		method:      method,
		handler:     handler,
		original:    handler,
		middlewares: make([]Middleware, 0),
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	route.Middleware(r.middlewares...)
	route.pattern = transform(route.pattern)
	if _, ok := r.routes[route.pattern]; ok {
		ErrorLog.Panicf("the \"%s\" route already exists", route.pattern)
	}
	r.routes[route.pattern] = route
	return route
}

func (r *Router) registerRoute(method string, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{
		pattern:     pattern,
		method:      method,
		handler:     handler,
		original:    handler,
		middlewares: make([]Middleware, 0),
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	route.Middleware(route.methodMiddleware)
	route.Middleware(r.middlewares...)
	route.pattern = transform(route.pattern)
	if _, ok := r.routes[route.pattern]; ok {
		ErrorLog.Panicf("the \"%s\" route already exists", route.pattern)
	}
	r.routes[route.pattern] = route
	return route
}
