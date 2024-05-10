package main

import (
	"encoding/json"
	"net/http"
)

func handleMutate(w http.ResponseWriter, r *http.Request) {
	var review AdmissionReview
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		log.Printf("Could not decode body: %v", err)
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	response := mutate(review)
	respondJSON(w, response)
}

func respondJSON(w http.ResponseWriter, response AdmissionResponse) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Could not encode response: %v", err)
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
