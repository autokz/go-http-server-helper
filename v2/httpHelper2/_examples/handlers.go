package _examples

import (
	"encoding/json"
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

func Welcome(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Welcome, " + r.RemoteAddr))
}

func Users(rw http.ResponseWriter, _ *http.Request) {
	users := []User{
		{
			ID:        1,
			Name:      "John Doe",
			CreatedAt: time.Now().UTC(),
		},
		{
			ID:        2,
			Name:      "Harold Lewis",
			CreatedAt: time.Now().UTC(),
		},
		{
			ID:        3,
			Name:      "Samia Levy",
			CreatedAt: time.Now().UTC(),
		},
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = rw.Write(b)
}

func Contacts(rw http.ResponseWriter, _ *http.Request) {
	contacts := []Contact{
		{
			UserID: 1,
			Phone:  "89991515555",
		},
		{
			UserID: 2,
			Phone:  "89991516666",
		},
		{
			UserID: 3,
			Phone:  "89991517777",
		},
	}

	b, err := json.Marshal(contacts)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = rw.Write(b)
}
