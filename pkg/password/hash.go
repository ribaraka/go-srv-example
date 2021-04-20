package password

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		// TODO: please handle this error
		log.Println(err)
	}

	return string(hash)
}