package httpHelper

import (
	"fmt"
	"net/http"
)

// getDefaultOptionsMiddleware
// Deprecated: The function is deprecated use the httpHelper2 package.
func getDefaultOptionsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == OPTIONS_METHOD {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Satrap-1, X-Satrap-2, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
			w.Header().Set("Access-Control-Expose-Headers", "X-Satrap-1, X-Satrap-2")

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("[]"))
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		next.ServeHTTP(w, r)
	}
}

// method
// Deprecated: The function is deprecated use the httpHelper2 package.
func method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)

			_, err := w.Write([]byte(`{"error":"` + r.Method + ` is not allowed for this route"}`))
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		next.ServeHTTP(w, r)
	}
}
