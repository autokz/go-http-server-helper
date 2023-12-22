package httpHelper

import (
	"net/http"
)

// Middleware represents a middleware function for HTTP handlers.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

var (
	JsonMiddleware = Middleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(writer, request)
		}
	})
	XmlMiddleware = Middleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "text/xml")
			next.ServeHTTP(writer, request)
		}
	})
	HtmlMiddleware = Middleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "text/html")
			next.ServeHTTP(writer, request)
		}
	})
	XhtmlMiddleware = Middleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/xhtml+xml")
			next.ServeHTTP(writer, request)
		}
	})
)
