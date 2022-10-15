package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lintopaul/go-projects/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
