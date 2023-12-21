package main

import (
	"github.com/autokz/go-http-server-helper/httpHelper2"
	"github.com/autokz/go-http-server-helper/httpHelper2/_examples"
	"log"
	"net/http"
)

func main() {
	router := httpHelper2.NewRouter()

	cors := httpHelper2.NewCORS()
	cors.AddAllowedHeaders("Authorization", "Content-Type")
	cors.SetAllowedMethods(http.MethodGet, http.MethodPost)
	cors.AddAllowedOrigins("http://localhost:3000", "http://localhost:3001")
	cors.SetExposeHeaders("Access-Token", "Refresh-Token")
	router.CORS(cors)

	router.Get("/", _examples.Welcome)

	log.Println(cors)
	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
