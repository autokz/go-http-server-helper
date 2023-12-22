package httpHelper2

import (
	"net/http"
)

// GroupRoute represents a group of routes within the router.
type GroupRoute struct {
	router      *Router
	pattern     string
	middlewares []Middleware
}

// NewGroupRoute creates a new subgroup route within the group.
func (gr *GroupRoute) NewGroupRoute(
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
	gr.router.mu.Lock()
	defer gr.router.mu.Unlock()
	if _, ok := gr.router.groups[gr.pattern+pattern]; ok {
		ErrorLog.Panicf("the \"%s\" group route already exists", gr.pattern+pattern)
	}
	ms := gr.middlewares
	if middlewares != nil {
		ms = append(ms, middlewares...)
	}
	groupRoute := &GroupRoute{
		router:      gr.router,
		pattern:     gr.pattern + pattern,
		middlewares: ms,
	}
	gr.router.groups[gr.pattern+pattern] = groupRoute
	return groupRoute
}

// GroupRoute creates a new subgroup route within the group and executes a callback function.
func (gr *GroupRoute) GroupRoute(pattern string, fn func(gr *GroupRoute), middlewares ...Middleware) *GroupRoute {
	if middlewares == nil {
		middlewares = make([]Middleware, 0)
	}

	if pattern == "" {
		panic("empty pattern")
	}
	if pattern[len(pattern)-1] == '/' {
		pattern = pattern[:len(pattern)-1]
	}

	gr.router.mu.Lock()
	if _, ok := gr.router.groups[gr.pattern+pattern]; ok {
		gr.router.mu.Unlock()
		ErrorLog.Panicf("the \"%s\" group route already exists", gr.pattern+pattern)
	}

	ms := gr.middlewares
	if middlewares != nil {
		ms = append(ms, middlewares...)
	}
	groupRoute := &GroupRoute{
		router:      gr.router,
		pattern:     gr.pattern + pattern,
		middlewares: ms,
	}
	gr.router.groups[gr.pattern+pattern] = groupRoute
	gr.router.mu.Unlock()
	fn(groupRoute)
	return groupRoute
}

// Middleware appends one or more middleware functions to the route.
func (gr *GroupRoute) Middleware(middlewares ...Middleware) *GroupRoute {
	if middlewares == nil {
		ErrorLog.Panic("empty middlewares")
	}
	gr.middlewares = append(gr.middlewares, middlewares...)
	return gr
}

// Get registers a new route with the GET method.
func (gr *GroupRoute) Get(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Get(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Head registers a new route with the Head method.
func (gr *GroupRoute) Head(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Head(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Post registers a new route with the Post method.
func (gr *GroupRoute) Post(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Post(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Put registers a new route with the Put method.
func (gr *GroupRoute) Put(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Put(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Patch registers a new route with the Patch method.
func (gr *GroupRoute) Patch(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Patch(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Delete registers a new route with the Delete method.
func (gr *GroupRoute) Delete(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Delete(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Connect registers a new route with the Connect method.
func (gr *GroupRoute) Connect(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Connect(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Options registers a new route with the Options method.
func (gr *GroupRoute) Options(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Options(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// Trace registers a new route with the Trace method.
func (gr *GroupRoute) Trace(pattern string, handler http.HandlerFunc) *Route {
	return gr.router.Trace(gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}

// HandleFunc registers a new route with a custom HTTP method.
func (gr *GroupRoute) HandleFunc(method string, pattern string, handler http.HandlerFunc) *Route {
	return gr.router.HandleFunc(method, gr.pattern+pattern, handler).Middleware(gr.middlewares...)
}
