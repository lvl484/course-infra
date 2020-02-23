package api

import "net/http"

func health(msg string) http.HandlerFunc {
	body := []byte(msg)

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}
}
