package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

type AuthResponse struct {
	Allowed    bool   `json:"allowed"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`
	InstanceID string `json:"instance_id"`
}

var instanceID = uuid.New().String()

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestID := uuid.New().String()

	// Randomly succeed or fail.
	allowed := rand.Intn(2) == 0

	resp := AuthResponse{
		Allowed:    allowed,
		Message:    "auth check complete",
		RequestID:  requestID,
		InstanceID: instanceID,
	}

	w.Header().Set("Content-Type", "application/json")

	if allowed {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/validate", validateHandler)

	log.Println("Auth server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
