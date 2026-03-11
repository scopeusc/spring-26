package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Response struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Message:   "Hello from Go 👋",
		RequestID: uuid.New().String(),
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
