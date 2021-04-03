package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string `json:"firstName" validate:"required,lte=5"`
	LastName  string `json:"lastName" validate:"required,lte=5"`
	Email     string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=64"`
}

func main() {
	form := http.FileServer(http.Dir("./form"))
	http.Handle("/", form)
	http.HandleFunc("/form", POSTHandler)
	log.Println("Listening...")
	v := validator.New()
	a := User{
		Email: "asd@afsaasdasdw2e",
	}
	errV := v.Struct(a)

	for _, e := range errV.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func POSTHandler(w http.ResponseWriter, r *http.Request) {

	db := OpenConnection()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", u)
	w.WriteHeader(201)
	sqlStatement := `INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	defer db.Close()
}

const (
	host     = "localhost"
	port     = 33333
	user     = "postgres"
	password = "password"
	dbname   = "go_project"
)

func OpenConnection() *sql.DB {
	PSQLInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", PSQLInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
