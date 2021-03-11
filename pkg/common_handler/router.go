package common_handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Group ...
func Group(router *mux.Router, path string) *mux.Router {
	return router.PathPrefix(path).Subrouter()
}

// GET ...
func GET(router *mux.Router, path string, handler http.HandlerFunc) {
	router.HandleFunc(path, handler).Methods(http.MethodGet)
}

// POST ...
func POST(router *mux.Router, path string, handler http.HandlerFunc) {
	router.HandleFunc(path, handler).Methods(http.MethodPost)
}
