package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPass(inputPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inputPass), []byte(dbPass))
	if err != nil {
		return err == nil
	}

	return true
}