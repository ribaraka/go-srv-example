package crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hash), nil
}


func CheckPassword(dbPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}