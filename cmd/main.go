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
	Password  string `json:"password" validate:"required,gte=8,lte=64"`
}

func main() {
	form := http.FileServer(http.Dir("./form"))
	http.Handle("/", form)
	http.HandleFunc("/form", POSTHandler)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", user)
	sqlStatement := `INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password)
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
