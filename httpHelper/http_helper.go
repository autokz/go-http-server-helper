package httpHelper

import (
	"net/http"
)

type Route struct {
	UrlPattern string
	Handler    http.Handler
}

type HandlerFunc func(next http.Handler) http.Handler

type Handler2Func func(w http.ResponseWriter, r *http.Request)

type MethodFunc func(method string, w http.ResponseWriter, r *http.Request)
