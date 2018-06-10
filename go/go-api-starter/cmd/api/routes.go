package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TestEndpoint will return the welcome message for test in JSON as an example.
func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("Test endpoint is being hit now!")

	// Call to some function to do some work.
	// If you get an error back from a function you can respond with a header
	// of some error to inform the user, e.g.
	// w.WriteHeader(http.StatusInternalServerError)

	// Build response if it was not received fom function.
	m := make(map[string]string)
	m["endpoint"] = "/v1/test"
	m["response"] = "You have hit the GET test endpoint and it has answered!"

	// Marshal response and write it back.
	respondJSON, _ := json.Marshal(m)
	w.Write(respondJSON)
}
