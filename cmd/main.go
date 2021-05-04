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

	pool, err := postgres.OpenConnection(config)
	if err != nil {
		log.Fatal("cannot initiate connection to database:", err)
	}

	defer pool.Close()

	signupRepo := postgres.NewSignUpRepository(pool)

	postHandler := handlers.NewPostHandler(signupRepo)


	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./form"))).Methods(http.MethodGet)
	r.HandleFunc("/form", postHandler).Methods(http.MethodPost)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(config.ServerLocalHost, r))
}