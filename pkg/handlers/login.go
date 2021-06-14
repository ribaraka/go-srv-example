package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"github.com/ribaraka/go-srv-example/pkg/auth"
	"github.com/ribaraka/go-srv-example/pkg/crypto"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/pkg/postgres"
	"log"
	"net/http"
)

func SignIn(l *postgres.LoginRepository, repo *postgres.SignUpRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Login
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
		requestUser, err := repo.GetByEmail(ctx, user.Email)
		if err != nil {
			if err == pgx.ErrNoRows {
				err := fmt.Errorf("no user found %v", err)
				log.Println(err)
				http.Error(w, err.Error(), 404)
				return
			}
			err := fmt.Errorf("emaildoes not exist  %v ", err)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ReqUserPassword, err := l.GetByID(ctx, requestUser.Id)
		if err != nil {
			if err == pgx.ErrNoRows {
				err := fmt.Errorf("no password found %v", err)
				log.Println(err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err := fmt.Errorf("pass not exist  %v ", err)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = crypto.CheckPassword(ReqUserPassword.PasswordHash, user.Password)
		if err != nil {
			err := fmt.Errorf("wrong passsword %s", err)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("correct password"))


		jwtWrapper := auth.JwtWrap{
			SecretKey:       "mySecretKey",
		}

		accessToken, err := jwtWrapper.GenerateAccessToken(user.Email,30)
		refreshToken, err := jwtWrapper.GenerateRefreshToken(user.Email,72)


		return
	}
}
