package httpHelper2

import (
	"net/http"
	"strings"
	"sync"
)

// Router represents a simple HTTP router with support for middleware and CORS.
type Router struct {
	mux         *http.ServeMux
	mu          *sync.Mutex
	routes      map[string]*Route
	groups      map[string]*GroupRoute
	middlewares []Middleware
	cors        *CORS
}

// NewRouter creates a new instance of the Router.
func NewRouter(middlewares ...Middleware) *Router {
	if middlewares == nil {
		middlewares = make([]Middleware, 0)
	}
	return &Router{
		mux:         http.NewServeMux(),
		mu:          &sync.Mutex{},
		routes:      make(map[string]*Route),
		groups:      make(map[string]*GroupRoute),
		middlewares: middlewares,
		cors:        nil,
	}
}

// NewGroupRoute creates a new group route within the router.
func (r *Router) NewGroupRoute(
	pattern string,
	middlewares ...Middleware,
) *GroupRoute {
	if middlewares == nil {
		middlewares = make([]Middleware, 0)
	}
	if pattern == "" {
		ErrorLog.Panic("empty pattern")
	}
	if pattern[len(pattern)-1] == '/' {
		pattern = pattern[:len(pattern)-1]
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.groups[pattern]; ok {
		ErrorLog.Panicf("the \"%s\" group route already exists", pattern)
	}
	gr := &GroupRoute{
		router:      r,
		pattern:     pattern,
		middlewares: middlewares,
	}
	r.groups[pattern] = gr
	return gr
}

// GroupRoute creates a new group route within the router and executes a callback function.
func (r *Router) GroupRoute(pattern string, fn func(gr *GroupRoute), middlewares ...Middleware) *GroupRoute {
	if middlewares == nil {
		middlewares = make([]Middleware, 0)
	}
	r.mu.Lock()
	pattern = strings.TrimSpace(pattern)
	if pattern == "" || pattern == "/" {
		r.mu.Unlock()
		ErrorLog.Panic("empty pattern")
	}
	if pattern[len(pattern)-1] == '/' {
		pattern = pattern[:len(pattern)-1]
	}

	if _, ok := r.groups[pattern]; ok {
		r.mu.Unlock()
		ErrorLog.Panicf("the \"%s\" group route already exists", pattern)
	}
	gr := &GroupRoute{
		router:      r,
		pattern:     pattern,
		middlewares: middlewares,
	}
	r.groups[pattern] = gr
	r.mu.Unlock()
	fn(gr)
	return gr
}

// Middleware appends one or more middleware functions to the route.
func (r *Router) Middleware(middlewares ...Middleware) *Router {
	if middlewares == nil {
		ErrorLog.Panic("empty middlewares")
	}
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

// CORS sets the CORS configuration for the router.
func (r *Router) CORS(cors *CORS) *Router {
	r.cors = cors
	return r
}

// Mux returns the underlying HTTP serve mux for the router.
func (r *Router) Mux() *http.ServeMux {
	r.handle()
	return r.mux
}

// handle configures the router's routes and their handlers.
func (r *Router) handle() *Router {
	for _, route := range r.routes {
		if r.cors != nil {
			route.handler = r.cors.Middleware(route.handler)
		}
		r.mux.HandleFunc(route.pattern, route.handler)
		InfoLog.Printf("%s %s", route.method, route.pattern)
	}
	return r
}
