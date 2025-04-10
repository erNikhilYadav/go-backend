package response

import (
	"encoding/json"
	"net/http"
	"time"
)

// StandardResponse represents the standard API response structure
type StandardResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// SuccessResponse creates a successful API response
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := StandardResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// ErrorResponse creates an error API response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := StandardResponse{
		Success:   false,
		Message:   message,
		Error:     err.Error(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
