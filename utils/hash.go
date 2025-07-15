package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	var pass = []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	if err != nil {
		log.Println("Unable to hash password.")
		return "", err
	}

	return string(hashedPassword), nil
}

func MatchPassword(password string, hashedPassword string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return check == nil
}
