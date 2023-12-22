package httpHelper2

import (
	"net/http"
	"strings"
)

// CORS represents the configuration for Cross-Origin Resource Sharing.
type CORS struct {
	allowedOrigins map[string]struct{}
	allowedMethods string
	allowedHeaders string
	exposeHeaders  string
}

// DefaultCORS is the default CORS configuration with some predefined values.
var DefaultCORS = &CORS{
	allowedOrigins: map[string]struct{}{
		"http://localhost:3000": {},
	},
	allowedMethods: strings.Join([]string{
		http.MethodHead,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	}, ", "),
	allowedHeaders: strings.Join([]string{
		"Authorization",
		"Access-Control-Allow-Headers",
		"Origin",
		"Accept",
		"X-Requested-With",
		"Content-Type",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}, ", "),
	exposeHeaders: "",
}

// NewCORS creates a new CORS configuration.
func NewCORS() *CORS {
	if DefaultCORS.allowedOrigins == nil {
		DefaultCORS.allowedOrigins = make(map[string]struct{})
	}
	return DefaultCORS
}

// Middleware is a middleware function that adds CORS headers to the response
func (cors *CORS) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for o := range cors.allowedOrigins {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", cors.allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", cors.allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", cors.exposeHeaders)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// AddAllowedOrigins adds allowed origins to the CORS configuration.
func (cors *CORS) AddAllowedOrigins(origins ...string) *CORS {
	for _, origin := range origins {
		cors.allowedOrigins[origin] = struct{}{}
	}
	return cors
}

// AddAllowedMethods adds allowed methods to the CORS configuration.
func (cors *CORS) AddAllowedMethods(methods ...string) *CORS {
	ms := make([]string, 0, len(methods))
	for _, method := range methods {
		if strings.Contains(cors.allowedMethods, method) {
			continue
		}
		ms = append(ms, method)
	}
	if strings.TrimSpace(cors.allowedMethods) != "" {
		cors.allowedMethods += ", " + strings.Join(ms, ", ")
		return cors
	}
	cors.allowedMethods += strings.Join(ms, ", ")
	return cors
}

// AddAllowedHeaders adds allowed headers to the CORS configuration.
func (cors *CORS) AddAllowedHeaders(headers ...string) *CORS {
	hs := make([]string, 0, len(headers))
	for _, header := range headers {
		if strings.Contains(cors.allowedHeaders, header) {
			continue
		}
		hs = append(hs, header)
	}
	if strings.TrimSpace(cors.allowedHeaders) != "" {
		cors.allowedHeaders += ", " + strings.Join(hs, ", ")
		return cors
	}
	cors.allowedHeaders += ", " + strings.Join(hs, ", ")
	return cors
}

// AddExposeHeaders adds expose headers to the CORS configuration.
func (cors *CORS) AddExposeHeaders(headers ...string) *CORS {
	hs := make([]string, 0, len(headers))
	for _, header := range headers {
		if strings.Contains(cors.exposeHeaders, header) {
			continue
		}
		hs = append(hs, header)
	}
	cors.exposeHeaders += strings.Join(hs, ", ")
	return cors
}

// SetAllowedOrigins sets allowed origins to the CORS configuration.
func (cors *CORS) SetAllowedOrigins(origins ...string) *CORS {
	for _, origin := range origins {
		cors.allowedOrigins[origin] = struct{}{}
	}
	return cors
}

// SetAllowedMethods sets allowed methods to the CORS configuration.
func (cors *CORS) SetAllowedMethods(methods ...string) *CORS {
	cors.allowedMethods = strings.Join(methods, ", ")
	return cors
}

// SetAllowedHeaders sets allowed headers to the CORS configuration.
func (cors *CORS) SetAllowedHeaders(headers ...string) *CORS {
	cors.allowedHeaders = strings.Join(headers, ", ")
	return cors
}

// SetExposeHeaders sets expose headers to the CORS configuration.
func (cors *CORS) SetExposeHeaders(headers ...string) *CORS {
	cors.exposeHeaders = strings.Join(headers, ", ")
	return cors
}
