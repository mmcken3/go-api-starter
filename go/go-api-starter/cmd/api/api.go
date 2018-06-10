package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Go API Starter is now running!")
	router := mux.NewRouter()

	// Configure all of the different API endpoints.
	router.HandleFunc("/v1/test", TestEndpoint).Methods("GET")

	// Handle the endpoints and serve on a set port.
	// This endpoint could be set in a env variable and pulled from there.
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8002", handler))
}
