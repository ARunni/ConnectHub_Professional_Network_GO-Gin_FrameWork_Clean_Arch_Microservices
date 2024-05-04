package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to generate password hash")
	}
	return string(hashPassword), nil
}
