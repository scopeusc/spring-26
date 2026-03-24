package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Response struct {
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`
	InstanceID string `json:"instance_id"`
}

var instanceID = uuid.New().String()

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Message:    "Hello from Go 👋",
		RequestID:  uuid.New().String(),
		InstanceID: instanceID,
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/google/uuid"
// )

// type Response struct {
// 	Message       string `json:"message"`
// 	RequestID     string `json:"request_id"`
// 	InstanceID    string `json:"instance_id"`
// 	AuthAllowed   bool   `json:"auth_allowed"`
// 	AuthMessage   string `json:"auth_message"`
// 	AuthRequestID string `json:"auth_request_id,omitempty"`
// }

// type AuthResponse struct {
// 	Allowed    bool   `json:"allowed"`
// 	Message    string `json:"message"`
// 	RequestID  string `json:"request_id"`
// 	InstanceID string `json:"instance_id"`
// }

// var instanceID = uuid.New().String()

// func authServiceURL() string {
// 	url := os.Getenv("AUTH_SERVICE_URL")
// 	if url == "" {
// 		return "http://auth-service:8081/validate"
// 	}
// 	return url
// }

// func callAuthService() (AuthResponse, error) {
// 	client := &http.Client{
// 		Timeout: 2 * time.Second,
// 	}

// 	resp, err := client.Get(authServiceURL())
// 	if err != nil {
// 		return AuthResponse{}, fmt.Errorf("call auth service: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	var authResp AuthResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
// 		return AuthResponse{}, fmt.Errorf("decode auth response: %w", err)
// 	}

// 	return authResp, nil
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	authResp, err := callAuthService()
// 	if err != nil {
// 		http.Error(w, "auth service unavailable", http.StatusBadGateway)
// 		return
// 	}

// 	if !authResp.Allowed {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusUnauthorized)

// 		_ = json.NewEncoder(w).Encode(Response{
// 			Message:       "Access denied by auth service",
// 			RequestID:     uuid.New().String(),
// 			InstanceID:    instanceID,
// 			AuthAllowed:   false,
// 			AuthMessage:   authResp.Message,
// 			AuthRequestID: authResp.RequestID,
// 		})
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	resp := Response{
// 		Message:       "Hello from Go 👋",
// 		RequestID:     uuid.New().String(),
// 		InstanceID:    instanceID,
// 		AuthAllowed:   true,
// 		AuthMessage:   authResp.Message,
// 		AuthRequestID: authResp.RequestID,
// 	}

// 	_ = json.NewEncoder(w).Encode(resp)
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)

// 	log.Println("Server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
//
