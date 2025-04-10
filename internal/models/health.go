package models

import "sync"

// HealthStatus represents the current health state
type HealthStatus struct {
	Status string `json:"status"`
	mu     sync.RWMutex
}

func (h *HealthStatus) GetStatus() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Status
}

func (h *HealthStatus) SetStatus(status string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Status = status
}
