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
	var confile =  flag.String("confile", "./cmd/config.yaml", "to specify config file please use -confile")
	flag.Parse()
	conf, err := config.LoadConfig(*confile)

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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(conf.StaticAssets))).Methods(http.MethodGet)
	r.HandleFunc("/form", postHandler).Methods(http.MethodPost)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(conf.ServerLocalHost, r))
}
