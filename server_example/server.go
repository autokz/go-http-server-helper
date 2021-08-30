package server

import (
	"github.com/autokz/go-http-server-helper/httpHelper"
	"github.com/autokz/go-http-server-helper/server_example/routes"
	"log"
	"net/http"
	"time"
)

func getRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	httpHelper.RegisterGroupRoutes(
		mux,
		&routes.TestRoutes,
		&routes.TestRoutes2,
	)

	httpHelper.RegisterRoute(mux, &routes.Route4)

	return mux
}

func InitServer() {
	server := &http.Server{
		Addr:         ":8099",
		Handler:      getRoutes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
