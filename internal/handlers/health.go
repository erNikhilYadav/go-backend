package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nikhilyadav/go-backend/internal/models"
)

type HealthHandler struct {
	health *models.HealthStatus
}

func NewHealthHandler(health *models.HealthStatus) *HealthHandler {
	return &HealthHandler{
		health: health,
	}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling GET /health request")
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": h.health.GetStatus()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *HealthHandler) ResetHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling POST /health/reset request")
	h.health.SetStatus("healthy")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status":  h.health.GetStatus(),
		"message": "Health status has been reset",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
