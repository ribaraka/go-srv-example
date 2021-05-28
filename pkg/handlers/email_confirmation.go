package handlers

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
	"log"
	"net/http"
)

func ConfirmEmail(tr *postgres.EmailTokenRepository, repo *postgres.SignUpRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		email := r.URL.Query().Get("email")
		if token == "" || email == "" {
			err := fmt.Errorf("please enter your email address and token in the URL for verification")
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		ctx := r.Context()
		requestUser, err := tr.GetByToken(ctx, token)
		if err != nil {
			if err == pgx.ErrNoRows{
				err := fmt.Errorf("no user token found %s", requestUser)
				log.Println(err)
				http.Error(w, err.Error(), 404)
				return
			}
			err := fmt.Errorf("cannot access to db %s", requestUser)
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		dbUser, err := repo.GetByID(ctx, requestUser.UserId)
		if err != nil {
			if err == pgx.ErrNoRows{
				err := fmt.Errorf("no user token found %v", requestUser.UserId)
				log.Println(err)
				http.Error(w, err.Error(), 404)
				return
			}
			err := fmt.Errorf("ID %v does not exist", dbUser)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if email != dbUser.Email {
			err := fmt.Errorf("cannot verify email")
			log.Println(err)
			http.Error(w, err.Error(), 404)
			return
		}

		fmt.Printf( "User %v Successfully Verified.", dbUser.Id)
		dbUser.Verified = true
		if err := repo.UpdateUserByEmail(ctx, dbUser); err != nil {
			er := fmt.Errorf("failed to update the validity %s", err)
			log.Println(er)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Successfully Verified."))
		return
	}
}
