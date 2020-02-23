package api

import (
	"encoding/json"
	"net/http"
)

func version(v string) http.HandlerFunc {
	body, _ := json.Marshal(struct {
		Version string `json:"version"`
	}{
		Version: v,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}
}
