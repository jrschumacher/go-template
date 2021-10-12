package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jrschumacher/go-template/internal/store"
)

// Update
// @Tags store
// @Summary Update data in store
// @Description Write data to store
// @Param id path string true "Id of data to be updated"
// @Param payload body map[string]interface{} true "Data to store"
// @Produce  json
// @Success 200 {object} store.StoreUpdateResult
// @Failure 400 {string} Bad Request
// @Failure 404 {string} Not Found
// @Failure 500 {string} Internal Server Error
// @Router /update/{id} [put]
func updateHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not find id", http.StatusNotFound)
		return
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Data is not a json object", http.StatusBadRequest)
		return
	}

	result, err := store.Update(id, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
