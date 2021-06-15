package handlers

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
	"log"
	"net/http"
	"strconv"
)

func GetProfile(repo *postgres.SignUpRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			err := fmt.Errorf("please enter your profile id in the URL")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		strId, err := strconv.Atoi(id)
		if err == nil {
			fmt.Println("id is not integer")
		}

		ctx := r.Context()
		requestUser, err := repo.GetByID(ctx, strId)
		if err != nil {
			if err == pgx.ErrNoRows {
				log.Println(err, "no user id found")
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			err := fmt.Errorf("cannot access to db %s", requestUser)
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}


		w.Write([]byte(requestUser.Firstname))
	}
}
