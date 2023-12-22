package main

import (
	"github.com/autokz/go-http-server-helper/v2/httpHelper"
	"github.com/autokz/go-http-server-helper/v2/httpHelper/_examples"
	"log"
	"net/http"
)

func main() {
	router := httpHelper.NewRouter()
	router.CORS(httpHelper.NewCORS())

	// Example 1
	v1Group := router.NewGroupRoute("/v1", httpHelper.JsonMiddleware)
	v1Group.Get("/users", _examples.Users)
	v1Group.Get("/users/contacts", _examples.Contacts)

	// Example 2
	v2Group := router.NewGroupRoute("/v2", httpHelper.JsonMiddleware)
	v2Users := v2Group.NewGroupRoute("/users", httpHelper.JsonMiddleware)
	v2Users.Get("", _examples.Users)
	v2Users.Get("/contacts", _examples.Users)

	// Example 3
	router.GroupRoute("/v3", func(gr *httpHelper.GroupRoute) {
		gr.Get("/users", _examples.Users)
		gr.Get("/users/contacts", _examples.Contacts)
	}, httpHelper.JsonMiddleware)

	// Example 4
	router.GroupRoute("/v4", func(gr *httpHelper.GroupRoute) {
		gr.Get("/users", _examples.Users)
		gr.GroupRoute("/users", func(gr *httpHelper.GroupRoute) {
			gr.Get("/contacts", _examples.Contacts)
		})
	}, httpHelper.JsonMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
