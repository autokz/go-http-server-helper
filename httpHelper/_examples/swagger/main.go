package main

import (
	"github.com/autokz/go-http-server-helper/v2/httpHelper"
	"github.com/autokz/go-http-server-helper/v2/httpHelper/_examples"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Contact struct {
	UserID int    `json:"user_id"`
	Phone  string `json:"phone"`
}

// @title 			Swagger Example
// @version			1.0
// @description		This example shows 3 ways to create a swagger documentation handler
// @BasePath		/
// @host			localhost:8080
func main() {
	router := httpHelper2.NewRouter(httpHelper2.JsonMiddleware)

	router.Get("/users", Users)
	router.Get("/users/contacts", Contacts)

	// Example 1
	router.Swagger("./docs") // http://localhost:8080/swagger

	// Example 2
	v1 := router.NewGroupRoute("/v1")
	v1.Swagger("./docs") // http://localhost:8080/v1/swagger

	// Example 3
	router.GroupRoute("/v2", func(gr *httpHelper2.GroupRoute) {
		gr.Swagger("./docs") // http://localhost:8080/v2/swagger
	})

	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}

// Users
// @Summary			 Get All Users.
// @Description      Returns a list of all users.
// @ID               get-all-users
// @Accept           json
// @Produce          json
// @Success          200                {object}   User    "Response object containing a list of all users"
// @Router           /users [get]
func Users(rw http.ResponseWriter, r *http.Request) {
	_examples.Users(rw, r)
	return
}

// Contacts
// @Summary			 Get User Contacts.
// @Description      Returns the contacts of all users.
// @ID               get-user-contacts
// @Accept           json
// @Produce          json
// @Success          200                {object}   Contact    "Response object containing a list of all the user's contacts"
// @Router           /users/contacts [get]
func Contacts(rw http.ResponseWriter, r *http.Request) {
	_examples.Contacts(rw, r)
	return
}
