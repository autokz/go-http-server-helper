package middleware

import (
	"log"
	"net/http"
)

// TEST
func M1(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(1)
		writer.Header().Set("M", "1")
		next(writer, request)
	}
}

func M2(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(2)
		writer.Header().Set("M2", "2")
		next(writer, request)
	}
}

func MainMiddleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("MainMiddleware1")
		writer.Header().Set("Main", "MAIN")
		next(writer, request)
	}
}

func MainMiddleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("MainMiddleware2")
		writer.Header().Set("Main2", "MAIN")
		next(writer, request)
	}
}
