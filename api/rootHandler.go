package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func rootHandler(router *mux.Router) (f func(w http.ResponseWriter, r *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		var results []map[string]interface{}

		if err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			r := make(map[string]interface{})
			r["path"], _ = route.GetPathTemplate()
			r["queries"], _ = route.GetQueriesTemplates()
			r["methods"], _ = route.GetMethods()
			results = append(results, r)
			return nil
		}); err != nil {
			logrus.WithError(err).Error("Could not create router map")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
