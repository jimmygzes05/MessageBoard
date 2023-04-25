package helper

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(input, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))

	return err == nil
}