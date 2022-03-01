package services

import (
	"daeng-market/queries"

	"golang.org/x/crypto/bcrypt"
)

func CreatePin(pin string, user_id int) (int, error) {
	hashed_pin, err := HashingPin(pin)
	if err != nil {
		return -1, err
	}
	pin_id, err := queries.CreatePinQueries(hashed_pin, user_id)
	if err != nil {
		return -1, err
	}
	return pin_id, nil
}

func HashingPin(pin string) (string, error) {
	token, err := bcrypt.GenerateFromPassword([]byte(pin), 14)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func CompareHashedPin(pin string, hashedPin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
	if err != nil {
		return false
	}
	return true
}
