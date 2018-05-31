// Package handlers provides the endpoints for the web service.

package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes sets the routes for the web service.
// It has 1 route call sendjson. When that route is executed, it will call the SendJSON function.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple JSON document.
// This has the same signature that we had before using ResponseWriter and Request.
// We create an anonymous struct, initialize it and unmarshall it into JSON and pass it down the
// line.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Hoanh An",
		Email: "hoanhan@bennington.edu",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
