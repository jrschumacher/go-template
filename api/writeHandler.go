package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/jrschumacher/go-template/internal/store"
)

const writeBadRequestMsg = "Data is not a json object"

// Write to Store
// @Tags store
// @Summary Write to store
// @Description Write data to store
// @Param payload body map[string]interface{} true "Data to store"
// @Produce  json
// @Success 200 {object} store.StoreSingleResult
// @Failure 400 {string} writeBadRequestMsg
// @Failure 500 {string} Internal Server Error
// @Router /write [post]
func writeHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, writeBadRequestMsg, http.StatusBadRequest)
		return
	}

	result, err := store.Add(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
