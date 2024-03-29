package main

import (
	"github.com/autokz/go-http-server-helper/v2/httpHelper"
	"github.com/autokz/go-http-server-helper/v2/httpHelper/_examples"
	"log"
	"net/http"
)

func main() {
	router := httpHelper2.NewRouter(httpHelper2.JsonMiddleware)

	cors := httpHelper2.NewCORS()
	cors.AddAllowedHeaders("Authorization", "Content-Type")
	cors.SetAllowedMethods(http.MethodGet, http.MethodPost)
	cors.AddAllowedOrigins("http://localhost:3000", "http://localhost:3001")
	cors.SetExposeHeaders("Access-Token", "Refresh-Token")
	router.CORS(cors)

	router.Get("/users", _examples.Users)

	log.Println(cors)
	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
