package httpHelper

import (
	"fmt"
	"net/http"
)

func Compose(endPointHandler Handler2Func, requestedMethod string, middlewares ...HandlerFunc) http.Handler {

	endPointHandlerWithMethod := method(requestedMethod, http.HandlerFunc(endPointHandler))
	countMiddleware := len(middlewares)

	if countMiddleware == 0 {
		return getOptionsMiddleware(endPointHandlerWithMethod)
	}

	var middlewareFunc HandlerFunc
	var middlewareChain http.Handler
	for i := 0; i < countMiddleware; i++ {

		middlewareFunc = middlewares[i]
		if i == 0 {
			middlewareChain = middlewareFunc(endPointHandlerWithMethod)
		} else {
			middlewareChain = middlewareFunc(middlewareChain)
		}
	}

	return getOptionsMiddleware(middlewareChain)
}

func getOptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

func method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)

			_, err := w.Write([]byte(`{"error":"` + r.Method + ` is not allowed for this route"}`))
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
