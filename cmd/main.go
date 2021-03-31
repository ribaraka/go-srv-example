package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func main() {
	form := http.FileServer(http.Dir("./form"))
	http.Handle("/", form)
	http.HandleFunc("/form", POSTHandler)
	log.Println("Listening...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
/*	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	fmt.Printf("%s\n", b)
	w.WriteHeader(201)*/

	//var u = [][]string{}
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	fmt.Printf("%+v\n", u)
	w.WriteHeader(201)
}