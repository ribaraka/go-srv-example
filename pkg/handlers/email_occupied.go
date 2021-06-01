package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
	"log"
	"net/http"
)

func CheckBusyEmail(repo *postgres.SignUpRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		ctx := r.Context()
		requestUser, err := repo.GetByEmail(ctx, user.Email)
		if err != nil {
			if err == pgx.ErrNoRows {
				return
			}
			err := fmt.Errorf("cannot access to db %s", requestUser)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("this email had been already registered"))
		return
	}
}