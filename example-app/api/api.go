package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func API(ver string, promExporter http.Handler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/health", health("OK")).Methods(http.MethodGet)
	r.HandleFunc("/version", version(ver)).Methods(http.MethodGet)
	r.Handle("/metrics", promExporter)
	r.HandleFunc("/example", example).Methods(http.MethodGet)

	return r
}
