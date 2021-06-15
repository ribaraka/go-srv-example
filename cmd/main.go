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
	var confile = flag.String("confile", "./cmd/config.yaml", "to specify config file please use flag -confile")
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
	postHandler := handlers.NewPostHandler(signUpRepo, conf)
	confirmEmailRepo := postgres.NewVerificationRepository(pool)
	getHandler := handlers.ConfirmEmail(confirmEmailRepo, signUpRepo)
	checkEmailRepo := postgres.NewSignUpRepository(pool)
	checkEmailHandler := handlers.CheckBusyEmail(checkEmailRepo)
	loginRepo := postgres.NewLoginRepository(pool)
	loginHandler := handlers.SignIn(loginRepo, signUpRepo)
	profileRepo := postgres.NewSignUpRepository(pool)
	profileHandler := handlers.GetProfile(profileRepo)


	r := mux.NewRouter()
	r.HandleFunc("/verify", getHandler)
	r.HandleFunc("/profile", profileHandler)
	r.HandleFunc("/possession", checkEmailHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/form", postHandler).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(conf.StaticAssets))).Methods(http.MethodGet)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(":8081", r))
}