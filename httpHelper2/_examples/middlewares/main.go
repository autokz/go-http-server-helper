package main

import (
	"github.com/autokz/go-http-server-helper/httpHelper2"
	"github.com/autokz/go-http-server-helper/httpHelper2/_examples"
	"log"
	"net/http"
)

func main() {
	// Example 1
	router := httpHelper2.NewRouter(m1, m2)
	router.Get("/", _examples.Welcome)

	// Example 2
	router.Middleware(m3, m4)
	router.Get("/welcome", _examples.Welcome)

	// Example 3
	router.Get("/users", _examples.Users).Middleware(m5, m6)

	// Example 4
	gr := router.NewGroupRoute("/v1/example", m7, m8)
	gr.Get("", _examples.Welcome)
	// Example 5
	gr.Middleware(m9, m10)

	// Example 6
	gr.GroupRoute("/welcome", func(gr *httpHelper2.GroupRoute) {
		gr.Get("", _examples.Contacts)
	}, m11, m12)

	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}

func m1(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m1")
		next.ServeHTTP(writer, request)
	}
}

func m2(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m2")
		next.ServeHTTP(writer, request)
	}
}

func m3(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m3")
		next.ServeHTTP(writer, request)
	}
}

func m4(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m4")
		next.ServeHTTP(writer, request)
	}
}

func m5(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m5")
		next.ServeHTTP(writer, request)
	}
}

func m6(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m6")
		next.ServeHTTP(writer, request)
	}
}

func m7(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m7")
		next.ServeHTTP(writer, request)
	}
}

func m8(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m8")
		next.ServeHTTP(writer, request)
	}
}

func m9(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m9")
		next.ServeHTTP(writer, request)
	}
}

func m10(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m10")
		next.ServeHTTP(writer, request)
	}
}

func m11(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m11")
		next.ServeHTTP(writer, request)
	}
}

func m12(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Middleware m12")
		next.ServeHTTP(writer, request)
	}
}
