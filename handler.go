package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler to manage the home endpoint(/)
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root endpoint working")
}

// Handler to manage the api endpoint
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API endpoint working")
}

// Receive a json and decode it
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

// Handler to manage the post request to user endpoint
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user Metadata
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", user)
}
