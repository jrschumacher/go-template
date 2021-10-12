package api

import (
	"encoding/json"
	"net/http"
)

type HealthResult struct {
	Ok bool `json:"ok"`
}

// Health
// @Tags utils
// @Summary Health endpoint
// @Description Get health of service
// @Produce  json
// @Success 200 {array} HealthResult
// @Router /health [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResult{Ok: true})
}
