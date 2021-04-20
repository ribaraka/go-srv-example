package main

import (
	"github.com/gorilla/mux"
	"github.com/ribaraka/go-srv-example/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./form"))).Methods(http.MethodGet)
	r.HandleFunc("/form", handlers.POSTHandler).Methods(http.MethodPost)
	log.Println("Server has been started...")
	// TODO: web server port should be configurable
	log.Fatal(http.ListenAndServe(":8081", r))
}
