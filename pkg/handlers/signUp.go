package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ribaraka/go-srv-example/config"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
)

func NewPostHandler(repo *postgres.SignUpRepository, c config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		v := validator.New()
		err = v.Struct(user)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				fmt.Println(e)
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		err = repo.AddUser(ctx, user, c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}