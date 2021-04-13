package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/postgres"
	"net/http"
)


func POSTHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user.Password)
	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := postgres.OpenConnection()
	sqlStatement := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, user.FirstName, user.LastName, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	defer db.Close()
}
