package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/i101dev/books-mysql/pkg/routes"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
