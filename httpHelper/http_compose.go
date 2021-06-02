package httpHelper

import (
	"fmt"
	"net/http"
)

func Compose(endPointHandler Handler2Func, requestedMethod string, middlewares ...HandlerFunc) http.Handler {

	endPointHandlerWithMethod := method(requestedMethod, http.HandlerFunc(endPointHandler))
	countMiddleware := len(middlewares)

	if countMiddleware == 0 {
		return endPointHandlerWithMethod
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

	return middlewareChain
}

func method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method && r.Method != OPTIONS_METHOD {
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
