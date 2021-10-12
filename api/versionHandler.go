package api

import (
	"encoding/json"
	"net/http"

	"github.com/jrschumacher/go-template/internal"
)

type VersionStat struct {
	Version     string `json:"version"`
	VersionLong string `json:"versionLong"`
	BuildTime   string `json:"buildTime"`
}

// Version
// @Tags utils
// @Summary Version endpoint
// @Description Get service version
// @Produce  json
// @Success 200 {object} VersionStat
// @Router /health [get]
func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(internal.GetVersion())
}
