package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 14)
	if error != nil {
		return "", errors.New("Cannot hash password" + error.Error())
	}

	return string(bytes), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
