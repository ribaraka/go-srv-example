package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ribaraka/go-srv-example/config"
	"github.com/ribaraka/go-srv-example/pkg/handlers"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
)

func main() {
	var path = flag.String("config", ".", "to specify path to config file please use -config")
	flag.Parse()
	log.Printf("config path is %s", *path)
	conf, err := config.LoadConfig(*path)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	pool, err := postgres.OpenConnection(conf)
	if err != nil {
		log.Fatal("cannot initiate connection to database:", err)
	}
	defer pool.Close()

	signUpRepo := postgres.NewSignUpRepository(pool)
	postHandler := handlers.NewPostHandler(signUpRepo)

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./form"))).Methods(http.MethodGet)
	r.HandleFunc("/form", postHandler).Methods(http.MethodPost)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(conf.ServerLocalHost, r))
}
