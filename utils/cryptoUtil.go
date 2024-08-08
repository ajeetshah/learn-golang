package utils

import "golang.org/x/crypto/bcrypt"

func GetPasswordHash(password string) string {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func IsPasswordMatch(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
