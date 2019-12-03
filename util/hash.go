package util


import (
	"golang.org/x/crypto/bcrypt"
)

//加密
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //bcrypt.MinCost
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

