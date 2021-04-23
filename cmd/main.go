package main

import (
	"github.com/gorilla/mux"
	"github.com/ribaraka/go-srv-example/handlers"
	"github.com/ribaraka/go-srv-example/internal/conf"
	"github.com/ribaraka/go-srv-example/postgres"
	"log"
	"net/http"
)

func main() {
	config, err := conf.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./form"))).Methods(http.MethodGet)
	r.HandleFunc("/form", handlers.POSTHandler).Methods(http.MethodPost)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(config.ServerLocalHost, r))
	postgres.OpenConnection()
}