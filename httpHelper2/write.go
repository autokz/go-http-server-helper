package httpHelper2

import (
	"encoding/json"
	"net/http"
)

// httpResponse represents a generic JSON response structure.
type httpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// writeJson writes a JSON response to the HTTP writer.
func writeJson(rw http.ResponseWriter, status int, s string) {
	r := httpResponse{
		Status:  status,
		Message: s,
	}
	b, err := json.Marshal(r)
	if err != nil {
		ErrorLog.Printf("json.Marshal(r) failed: %v | r - %+v", err, r)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, _ = rw.Write(b)
}
