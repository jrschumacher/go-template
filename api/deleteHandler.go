package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
)

// Delete
// @Tags store
// @Summary Delete an item
// @Description Delete an item from the store
// @Param id path string true "Id of data to be updated"
// @Produce  json
// @Success 200 {object} store.StoreDeleteResult
// @Failure 404 {string} Not Found
// @Router /delete/{id} [delete]
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := store.Delete(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Could not find id")
		http.Error(w, "Could not find id", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
