package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Cmp_Passwd(db string, user string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(db), []byte(user))
	return err == nil
}
