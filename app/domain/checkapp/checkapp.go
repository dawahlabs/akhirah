package checkapp

import (
	"encoding/json"
	"net/http"
)

func liveness(w http.ResponseWriter, h *http.Request) {
	status := struct {
		Message string
		Status  int
	}{
		Message: "liveness",
		Status:  http.StatusOK,
	}

	json.NewEncoder(w).Encode(status)
}

func readiness(w http.ResponseWriter, h *http.Request) {
	status := struct {
		Message string
		Status  int
	}{
		Message: "readiness",
		Status:  http.StatusOK,
	}

	json.NewEncoder(w).Encode(status)
}
