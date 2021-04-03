package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-playground/validator/v10"
)

func main() {
	form := http.FileServer(http.Dir("./form"))
	http.Handle("/", form)
	http.HandleFunc("/form", POSTHandler)
	log.Println("Server has been started...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
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
		return
	}

	db := OpenConnection()
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
