package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func buildRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).
		Methods("GET")

	return r
}
