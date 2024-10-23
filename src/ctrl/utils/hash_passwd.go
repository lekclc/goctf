package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash_passwd(data string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
