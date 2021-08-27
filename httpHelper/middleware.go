package httpHelper

import (
	"net/http"
)

type Middlewares []Middleware

type Middleware func(next http.HandlerFunc) http.HandlerFunc
