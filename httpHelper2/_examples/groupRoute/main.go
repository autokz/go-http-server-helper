package main

import (
	"github.com/autokz/go-http-server-helper/httpHelper2"
	"github.com/autokz/go-http-server-helper/httpHelper2/_examples"
	"log"
	"net/http"
)

func main() {
	router := httpHelper2.NewRouter()
	router.CORS(httpHelper2.NewCORS())

	router.Get("/", _examples.Welcome)

	// Example 1
	v1Group := router.NewGroupRoute("/v1", httpHelper2.JsonMiddleware)
	v1Group.Get("/users", _examples.Users)
	v1Group.Get("/users/contacts", _examples.Contacts)

	// Example 2
	v2Group := router.NewGroupRoute("/v2", httpHelper2.JsonMiddleware)
	v2Users := v2Group.NewGroupRoute("/users", httpHelper2.JsonMiddleware)
	v2Users.Get("", _examples.Users)
	v2Users.Get("/contacts", _examples.Users)

	// Example 3
	router.GroupRoute("/v3", func(gr *httpHelper2.GroupRoute) {
		gr.Get("/users", _examples.Users)
		gr.Get("/users/contacts", _examples.Contacts)
	}, httpHelper2.JsonMiddleware)

	// Example 4
	router.GroupRoute("/v4", func(gr *httpHelper2.GroupRoute) {
		gr.Get("/users", _examples.Users)
		gr.GroupRoute("/users", func(gr *httpHelper2.GroupRoute) {
			gr.Get("/contacts", _examples.Contacts)
		})
	}, httpHelper2.JsonMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
