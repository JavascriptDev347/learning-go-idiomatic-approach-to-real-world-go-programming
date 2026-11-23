package stock

import (
	"encoding/json"
	"net/http"
	"task_context/internal/tracker"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tracker.Log(r.Context()).Info("stock requested", "item", r.PathValue("id"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": 7})
}
