package api

import (
	"encoding/json"
	"net/http"

	_ "github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
)

// Search
// @Tags store
// @Summary Search for data in store
// @Description Search for data in store
// @Produce  json
// @Success 200 {object} store.StoreMultiResult
// @Failure 500 {string} Internal Server Error
// @Router /search [get]
func searchHandler(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		params[k] = v
	}

	logrus.WithFields(params).Info("Search query")
	data, err := store.Find(params)
	if err != nil {
		logrus.WithError(err).WithFields(params).Error("Could not find data")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
