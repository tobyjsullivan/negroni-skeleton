package main

import (
	"os"

	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := buildRoutes()

	n := negroni.New()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)
}

func buildRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).
		Methods("GET")

	return r
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}
